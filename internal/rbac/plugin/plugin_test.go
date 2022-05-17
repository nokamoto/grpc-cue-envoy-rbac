package plugin

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api/api"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestPlugin_Run(t *testing.T) {
	extend := func(v *api.Authorization) *descriptorpb.MethodOptions {
		var opt descriptorpb.MethodOptions
		proto.SetExtension(&opt, api.E_Authz, v)
		return &opt
	}

	foo := descriptorpb.FileDescriptorProto{
		Name:    proto.String("foo"),
		Package: proto.String("bar"),
		Service: []*descriptorpb.ServiceDescriptorProto{
			{
				Name: proto.String("Baz"),
				Method: []*descriptorpb.MethodDescriptorProto{
					{
						Name: proto.String("Qux"),
						Options: extend(&api.Authorization{
							Permission: "quux",
						}),
					},
				},
			},
		},
	}

	tests := []struct {
		name string
		req  *pluginpb.CodeGeneratorRequest
		want *api.ExternalAuthorization
	}{
		{
			name: "empty request should be empty content",
			req:  &pluginpb.CodeGeneratorRequest{},
			want: &api.ExternalAuthorization{},
		},
		{
			name: "ExternalAuthorization_Rule should be added",
			req: &pluginpb.CodeGeneratorRequest{
				FileToGenerate: []string{"foo"},
				ProtoFile:      []*descriptorpb.FileDescriptorProto{&foo},
			},
			want: &api.ExternalAuthorization{
				Rules: []*api.ExternalAuthorization_Rule{
					{
						Path: "/bar.Baz/Qux",
						Authorization: &api.Authorization{
							Permission: "quux",
						},
					},
				},
			},
		},
		{
			name: "ProtoFile should be excluded if not in FileToGenerate",
			req: &pluginpb.CodeGeneratorRequest{
				ProtoFile: []*descriptorpb.FileDescriptorProto{&foo},
			},
			want: &api.ExternalAuthorization{},
		},
		{
			name: "Method without Authorization should be excluded",
			req: &pluginpb.CodeGeneratorRequest{
				FileToGenerate: []string{"foo"},
				ProtoFile: []*descriptorpb.FileDescriptorProto{
					{
						Name:    proto.String("foo"),
						Package: proto.String("bar"),
						Service: []*descriptorpb.ServiceDescriptorProto{
							{
								Name: proto.String("Baz"),
								Method: []*descriptorpb.MethodDescriptorProto{
									{
										Name: proto.String("Qux"),
									},
								},
							},
						},
					},
				},
			},
			want: &api.ExternalAuthorization{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, _ := proto.Marshal(test.req)

			var buf bytes.Buffer

			p := Plugin{
				input:  bytes.NewReader(req),
				output: &buf,
			}

			if err := p.Run(); err != nil {
				t.Fatal(err)
			}

			var got pluginpb.CodeGeneratorResponse
			if err := proto.Unmarshal(buf.Bytes(), &got); err != nil {
				t.Fatal(err)
			}

			content, err := protojson.MarshalOptions{
				Indent: "  ",
			}.Marshal(test.want)
			if err != nil {
				t.Fatal(err)
			}

			want := pluginpb.CodeGeneratorResponse{
				File: []*pluginpb.CodeGeneratorResponse_File{
					{
						Name:    proto.String("rbac.json"),
						Content: proto.String(string(content) + "\n"),
					},
				},
			}

			if diff := cmp.Diff(&want, &got, protocmp.Transform()); diff != "" {
				t.Errorf("unexpected response (-want +got):\n%s", diff)
			}
		})
	}
}

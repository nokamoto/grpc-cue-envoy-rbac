package plugin

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestPlugin_Run(t *testing.T) {
	req, _ := proto.Marshal(&pluginpb.CodeGeneratorRequest{})

	var buf bytes.Buffer

	p := Plugin{
		input:  bytes.NewReader(req),
		output: &buf,
	}

	if err := p.Run(); err != nil {
		t.Fatal(err)
	}

	var actual pluginpb.CodeGeneratorResponse
	if err := proto.Unmarshal(buf.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}

	expected := pluginpb.CodeGeneratorResponse{
		File: []*pluginpb.CodeGeneratorResponse_File{
			{
				Name:    proto.String("rbac.json"),
				Content: proto.String("{}\n"),
			},
		},
	}

	if diff := cmp.Diff(&expected, &actual, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected response (-want +got):\n%s", diff)
	}
}

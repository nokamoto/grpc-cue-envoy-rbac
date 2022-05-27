package authorization

import (
	"context"
	"testing"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/service/authorization/mock"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"
)

func newRequest(path string, headers map[string]string) *v3.CheckRequest {
	return &v3.CheckRequest{
		Attributes: &v3.AttributeContext{
			Request: &v3.AttributeContext_Request{
				Http: &v3.AttributeContext_HttpRequest{
					Path:    path,
					Headers: headers,
				},
			},
		},
	}
}

func TestAuthorization_Check(t *testing.T) {
	rules := []*api.ExternalAuthorization_Rule{
		{
			Path: "/ok",
			Authorization: &api.Authorization{
				Permission: "foo",
			},
		},
	}

	tests := []struct {
		name string
		mock func(rbac *mock.MockRBAC)
		req  *v3.CheckRequest
		want *v3.CheckResponse
		code codes.Code
	}{
		{
			name: "allow reflection",
			req:  newRequest("/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo", nil),
			want: newResponse(codes.OK, ""),
		},
		{
			name: "allow unknown path",
			req:  newRequest("/unknown", nil),
			want: newResponse(codes.OK, ""),
		},
		{
			name: "allow if authorization passed",
			mock: func(rbac *mock.MockRBAC) {
				rbac.EXPECT().AuthorizeUser(gomock.Any(), &api.AuthorizeUserRequest{
					User: "bar",
					Authorization: &api.Authorization{
						Permission: "foo",
					},
				}).Return(nil, nil)
			},
			req:  newRequest("/ok", map[string]string{"x-username": "bar"}),
			want: newResponse(codes.OK, ""),
		},
		{
			name: "deny if authorization failed",
			mock: func(rbac *mock.MockRBAC) {
				rbac.EXPECT().AuthorizeUser(gomock.Any(), &api.AuthorizeUserRequest{
					User: "bar",
					Authorization: &api.Authorization{
						Permission: "foo",
					},
				}).Return(nil, status.Error(codes.PermissionDenied, ""))
			},
			req:  newRequest("/ok", map[string]string{"x-username": "bar"}),
			want: newResponse(codes.PermissionDenied, "permission denied"),
		},
		{
			name: "deny if unexpected error",
			mock: func(rbac *mock.MockRBAC) {
				rbac.EXPECT().AuthorizeUser(gomock.Any(), &api.AuthorizeUserRequest{
					User: "bar",
					Authorization: &api.Authorization{
						Permission: "foo",
					},
				}).Return(nil, status.Error(codes.DeadlineExceeded, ""))
			},
			req:  newRequest("/ok", map[string]string{"x-username": "bar"}),
			want: newResponse(codes.Internal, "internal error occurs"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			rbac := mock.NewMockRBAC(ctrl)
			if test.mock != nil {
				test.mock(rbac)
			}

			a := &Authorization{
				cfg: &api.ExternalAuthorization{
					Rules: rules,
				},
				rbac: rbac,
			}
			got, err := a.Check(context.TODO(), test.req)
			if status.Code(err) != test.code {
				t.Errorf("Check() error = %v, want %v", err, test.code)
			}
			if diff := cmp.Diff(test.want, got, protocmp.Transform()); len(diff) != 0 {
				t.Errorf("Check() diff = %v", diff)
			}
		})
	}
}

package authorization

import (
	"context"
	"testing"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"
)

func newRequest(path string) *v3.CheckRequest {
	return &v3.CheckRequest{
		Attributes: &v3.AttributeContext{
			Request: &v3.AttributeContext_Request{
				Http: &v3.AttributeContext_HttpRequest{
					Path: path,
				},
			},
		},
	}
}

func TestAuthorization_Check(t *testing.T) {
	tests := []struct {
		name string
		req  *v3.CheckRequest
		want *v3.CheckResponse
		code codes.Code
	}{
		{
			name: "allow reflection",
			req:  newRequest("/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"),
			want: newResponse(codes.OK, ""),
		},
		{
			name: "permission dinied",
			req:  newRequest("/unknown"),
			want: newResponse(codes.PermissionDenied, "permission denied"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := &Authorization{}
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

package authorization

import (
	"context"
	"log"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/rbac/plugin"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

// Authorization implements v3.AuthorizationServer.
type Authorization struct {
	cfg *plugin.Config
	v3.UnimplementedAuthorizationServer
}

func NewAuthorization(cfg *plugin.Config) *Authorization {
	return &Authorization{cfg: cfg}
}

const (
	reflectionPath = "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
)

func newResponse(code codes.Code, message string) *v3.CheckResponse {
	return &v3.CheckResponse{
		Status: &status.Status{
			Code:    int32(code),
			Message: message,
		},
	}
}

func (*Authorization) Check(ctx context.Context, req *v3.CheckRequest) (*v3.CheckResponse, error) {
	path := req.GetAttributes().GetRequest().GetHttp().GetPath()
	log.Printf("Check: %v", path)
	if path == reflectionPath {
		return newResponse(codes.OK, ""), nil
	}
	return newResponse(codes.PermissionDenied, "permission denied"), nil
}

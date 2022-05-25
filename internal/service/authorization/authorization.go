package authorization

import (
	"context"
	"log"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/rbac/plugin"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	googleapisstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Authorization implements v3.AuthorizationServer.
type Authorization struct {
	cfg  *plugin.Config
	rbac RBAC
	v3.UnimplementedAuthorizationServer
}

func NewAuthorization(cfg *plugin.Config, rbac RBAC) *Authorization {
	return &Authorization{cfg: cfg, rbac: rbac}
}

const (
	reflectionPath = "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
)

func newResponse(code codes.Code, message string) *v3.CheckResponse {
	return &v3.CheckResponse{
		Status: &googleapisstatus.Status{
			Code:    int32(code),
			Message: message,
		},
	}
}

func (a *Authorization) match(path string) *api.ExternalAuthorization_Rule {
	for _, rule := range a.cfg.GetRules() {
		if rule.GetPath() == path {
			return rule
		}
	}
	return nil
}

func (a *Authorization) authorize(ctx context.Context, req *v3.CheckRequest, rule *api.ExternalAuthorization_Rule) *v3.CheckResponse {
	if rule == nil {
		return newResponse(codes.OK, "")
	}
	_, err := a.rbac.AuthorizeUser(ctx, &api.AuthorizeUserRequest{
		User:          "todo",
		Authorization: rule.GetAuthorization(),
	})
	switch status.Code(err) {
	case codes.OK:
		return newResponse(codes.OK, "")
	case codes.PermissionDenied:
		return newResponse(codes.PermissionDenied, "permission denied")
	}
	log.Printf("AuthorizeUser: %v", err)
	return newResponse(codes.Internal, "internal error occurs")
}

func (a *Authorization) Check(ctx context.Context, req *v3.CheckRequest) (*v3.CheckResponse, error) {
	path := req.GetAttributes().GetRequest().GetHttp().GetPath()
	log.Printf("Check: %v", path)
	if path == reflectionPath {
		return newResponse(codes.OK, ""), nil
	}
	return a.authorize(ctx, req, a.match(path)), nil
}

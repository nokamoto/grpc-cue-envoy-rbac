//go:generate mockgen -destination mock/mock_rbac.go -package mock . RBAC
package authorization

import (
	"context"

	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/grpc"
)

// RBAC represents api.RBACServiceClient.
type RBAC interface {
	AuthorizeUser(ctx context.Context, in *api.AuthorizeUserRequest, opts ...grpc.CallOption) (*api.AuthorizeUserResponse, error)
}

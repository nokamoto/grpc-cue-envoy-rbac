package rbac

import (
	"context"
	"log"

	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RBAC implements api.RBACServiceServer.
type RBAC struct {
	table Table
	api.UnimplementedRBACServiceServer
}

// Username is a user name string.
type Username = string

// Permissions is a list of a permission string.
type Permissions = []string

// Table is an allow list.
//
// A user is allowed to access a resource if the user has a requested permission.
type Table = map[Username]Permissions

// NewRBAC returns a new RBAC service.
func NewRBAC(table Table) *RBAC {
	return &RBAC{table: table}
}

// AuthorizeUser returns OK if the user has a requested permission, PermissionDenied otherwise.
func (r *RBAC) AuthorizeUser(ctx context.Context, req *api.AuthorizeUserRequest) (*api.AuthorizeUserResponse, error) {
	log.Printf("AuthorizeUser: %v", req)
	permissions := r.table[req.GetUser()]
	for _, permission := range permissions {
		if permission == req.GetAuthorization().GetPermission() {
			return &api.AuthorizeUserResponse{}, nil
		}
	}
	return nil, status.Errorf(codes.PermissionDenied, "user '%s' does not have permission '%s'", req.GetUser(), req.GetAuthorization().GetPermission())
}

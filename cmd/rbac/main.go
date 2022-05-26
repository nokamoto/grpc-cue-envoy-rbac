package main

import (
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/server"
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/service/rbac"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	server.Serve(9002, func(s *grpc.Server) error {
		api.RegisterRBACServiceServer(s, rbac.NewRBAC(rbac.Table{
			"nokamoto@example.com": []string{"example.example.create"},
		}))
		return nil
	})
}

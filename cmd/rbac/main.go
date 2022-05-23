package main

import (
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/server"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	server.Serve(9002, func(s *grpc.Server) error {
		api.RegisterRBACServiceServer(s, &api.UnimplementedRBACServiceServer{})
		return nil
	})
}

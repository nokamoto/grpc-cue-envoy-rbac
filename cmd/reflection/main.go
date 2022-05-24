package main

import (
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/server"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	server.Serve(9003, func(s *grpc.Server) error {
		// register all unimplemented services for refrection.
		// ref. https://github.com/grpc/grpc-go/tree/master/reflection
		api.RegisterExampleServiceServer(s, &api.UnimplementedExampleServiceServer{})
		api.RegisterRBACServiceServer(s, &api.UnimplementedRBACServiceServer{})
		return nil
	})
}

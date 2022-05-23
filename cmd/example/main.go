package main

import (
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/server"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/grpc"
)

type example struct {
	api.UnimplementedExampleServiceServer
}

func main() {
	server.Serve(9000, func(s *grpc.Server) error {
		api.RegisterExampleServiceServer(s, &example{})
		return nil
	})
}

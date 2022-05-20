package main

import (
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/server"
	"google.golang.org/grpc"
)

type authorization struct {
	v3.UnimplementedAuthorizationServer
}

func main() {
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/examples/ext_authz/auth/grpc-service/main.go
	server.Serve(9001, func(s *grpc.Server) {
		v3.RegisterAuthorizationServer(s, &authorization{})
	})
}

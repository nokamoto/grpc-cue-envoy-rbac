package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type example struct {
	api.UnimplementedExampleServiceServer
}

func main() {
	port := 9000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterExampleServiceServer(s, &example{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

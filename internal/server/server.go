package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Serve starts a gRPC server with a reflection service.
//
// It calls os.Exit(1) if it fails to start a gRPC server.
func Serve(port int, f func(*grpc.Server) error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	if err := f(s); err != nil {
		log.Fatalf("failed to create: %v", err)
	}
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

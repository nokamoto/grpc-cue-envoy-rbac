package main

import (
	"fmt"
	"io/ioutil"
	"log"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/rbac/plugin"
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/server"
	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/service/authorization"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const configurationFile = "/etc/authorization/rbac.json"

func main() {
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/examples/ext_authz/auth/grpc-service/main.go
	server.Serve(9001, func(s *grpc.Server) error {
		data, err := ioutil.ReadFile(configurationFile)
		if err != nil {
			return err
		}

		cfg, err := plugin.Unmarshal(data)
		if err != nil {
			return err
		}

		log.Printf("config: %v", cfg.String())

		conn, err := grpc.Dial("rbac:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("did not connect: %v", err)
		}

		v3.RegisterAuthorizationServer(s, authorization.NewAuthorization(cfg, api.NewRBACServiceClient(conn)))
		return nil
	})
}

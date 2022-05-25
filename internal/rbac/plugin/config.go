package plugin

import (
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/protobuf/encoding/protojson"
)

// Config represents a plugin output file.
type Config = api.ExternalAuthorization

// Marshal returns JSON bytes of the configuration.
func Marshal(c *Config) ([]byte, error) {
	return protojson.Marshal(c)
}

// Unmarshal parses JSON bytes to the configuration.
func Unmarshal(data []byte) (*Config, error) {
	var c api.ExternalAuthorization
	if err := protojson.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

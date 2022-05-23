package plugin

import (
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/protobuf/encoding/protojson"
)

// Config represents a plugin output file.
type Config struct {
	api.ExternalAuthorization
}

// Marshal returns JSON bytes of the configuration.
func (c *Config) Marshal() ([]byte, error) {
	return protojson.Marshal(&c.ExternalAuthorization)
}

// Unmarshal parses JSON bytes to the configuration.
func (c *Config) Unmarshal(data []byte) error {
	return protojson.Unmarshal(data, &c.ExternalAuthorization)
}

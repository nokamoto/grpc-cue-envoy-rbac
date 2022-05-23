package plugin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestConfig(t *testing.T) {
	want := Config{
		ExternalAuthorization: api.ExternalAuthorization{
			Rules: []*api.ExternalAuthorization_Rule{
				{
					Path: "/foo",
					Authorization: &api.Authorization{
						Permission: "bar",
					},
				},
			},
		},
	}

	json, err := want.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	var got Config
	if err := got.Unmarshal(json); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(&want.ExternalAuthorization, &got.ExternalAuthorization, protocmp.Transform()); len(diff) != 0 {
		t.Errorf("diff: %s", diff)
	}
}

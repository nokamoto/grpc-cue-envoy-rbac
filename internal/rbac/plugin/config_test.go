package plugin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestConfig(t *testing.T) {
	want := &Config{
		Rules: []*api.ExternalAuthorization_Rule{
			{
				Path: "/foo",
				Authorization: &api.Authorization{
					Permission: "bar",
				},
			},
		},
	}

	json, err := Marshal(want)
	if err != nil {
		t.Fatal(err)
	}

	got, err := Unmarshal(json)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got, protocmp.Transform()); len(diff) != 0 {
		t.Errorf("diff: %s", diff)
	}
}

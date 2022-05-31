package plugin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api"
	"google.golang.org/protobuf/proto"
)

func FuzzConfig(f *testing.F) {
	// https://cs.opensource.google/go/go/+/master:src/encoding/json/fuzz_test.go
	bytes, _ := proto.Marshal(&api.ExternalAuthorization{
		Rules: []*api.ExternalAuthorization_Rule{
			{
				Path: "/foo",
				Authorization: &api.Authorization{
					Permission: "bar",
				},
			},
		},
	})
	f.Add(bytes)
	f.Fuzz(func(t *testing.T, want []byte) {
		c, err := Unmarshal(want)
		if err != nil {
			// skip
			return
		}

		t.Logf("data = %v", c)

		got, err := Marshal(c)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); len(diff) != 0 {
			t.Errorf("diff: %s", diff)
		}
	})
}

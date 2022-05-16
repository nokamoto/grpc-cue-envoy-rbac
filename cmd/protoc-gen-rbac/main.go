package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nokamoto/grpc-cue-envoy-rbac/internal/rbac/plugin"
)

func main() {
	if err := plugin.NewPlugin().Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}

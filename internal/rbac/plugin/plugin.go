package plugin

import (
	"io"
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

type Plugin struct {
	input  io.Reader
	output io.Writer
}

// NewPlugin creates a new Plugin.
//
// It reads a CodeGeneratorRequest message from os.Stdin, invokes the pulugin function, and writes a CodeGeneratorResponse message to os.Stdout.
func NewPlugin() *Plugin {
	return &Plugin{
		input:  os.Stdin,
		output: os.Stdout,
	}
}

// Run executes a function as a protoc plugin.
func (p *Plugin) Run() error {
	// ref. https://github.com/protocolbuffers/protobuf-go/blob/8a7ba0762cb3b39fc0536379eac2f7fa5796f187/compiler/protogen/protogen.go#L44-L56
	in, err := ioutil.ReadAll(p.input)
	if err != nil {
		return err
	}

	var req pluginpb.CodeGeneratorRequest
	if err := proto.Unmarshal(in, &req); err != nil {
		return err
	}

	out, err := proto.Marshal(&pluginpb.CodeGeneratorResponse{
		File: []*pluginpb.CodeGeneratorResponse_File{
			{
				Name:    proto.String("rbac.json"),
				Content: proto.String("{}"),
			},
		},
	})
	if err != nil {
		return err
	}

	if _, err := p.output.Write(out); err != nil {
		return err
	}
	return nil
}

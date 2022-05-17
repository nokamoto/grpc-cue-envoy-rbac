package plugin

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nokamoto/grpc-cue-envoy-rbac/pkg/api/api"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Plugin struct {
	input  io.Reader
	output io.Writer
	debug  io.Writer
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
//
// It generates a ExternalAuthorization message and writes it to `rbac.json`.
//
// `--rbac_opt=debug=stderr` enables debug output to stderr.
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

	p.parseParams(&req)

	var cfg api.ExternalAuthorization
	for _, file := range req.GetProtoFile() {
		p.debugf("file: %s\n", file.GetName())

		for _, s := range req.GetFileToGenerate() {
			if s == file.GetName() {
				if err := p.useFile(file, &cfg); err != nil {
					return err
				}
				break
			}
		}
	}

	content, err := protojson.MarshalOptions{
		Indent: "  ",
	}.Marshal(&cfg)
	if err != nil {
		return err
	}

	out, err := proto.Marshal(&pluginpb.CodeGeneratorResponse{
		File: []*pluginpb.CodeGeneratorResponse_File{
			{
				Name:    proto.String("rbac.json"),
				Content: proto.String(string(content) + "\n"),
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

func (p *Plugin) useFile(file *descriptorpb.FileDescriptorProto, cfg *api.ExternalAuthorization) error {
	for _, service := range file.GetService() {
		p.debugf("	service: %s\n", service.GetName())
		for _, method := range service.GetMethod() {
			p.debugf("		method: %s\n", method.GetName())
			if proto.HasExtension(method.GetOptions(), api.E_Authz) {
				authorization := proto.GetExtension(method.GetOptions(), api.E_Authz).(*api.Authorization)
				p.debugf("			authorization: %s\n", authorization)
				cfg.Rules = append(cfg.Rules, &api.ExternalAuthorization_Rule{
					Path:          fmt.Sprintf("/%s.%s/%s", file.GetPackage(), service.GetName(), method.GetName()),
					Authorization: authorization,
				})
			}
		}
	}
	return nil
}

func (p *Plugin) parseParams(req *pluginpb.CodeGeneratorRequest) {
	for _, param := range strings.Split(req.GetParameter(), ",") {
		var value string
		if i := strings.Index(param, "="); i >= 0 {
			value = param[i+1:]
			param = param[0:i]
		}
		switch param {
		case "debug":
			switch value {
			case "stderr":
				p.debug = os.Stderr
			}
			p.debugf("debug=%s\n", value)
		default:
			// ignore unknown parameters
		}
	}
}

func (p *Plugin) debugf(format string, a ...any) {
	if p.debug == nil {
		return
	}
	fmt.Fprintf(p.debug, format, a...)
}

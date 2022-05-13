GO_OUT = pkg/api

.PHONY: all
all:
	find api -name *.proto | xargs clang-format -i
	go fmt ./...

	protoc --go_out=$(GO_OUT) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT) \
		--go-grpc_opt=paths=source_relative \
		-I api \
		$$(find api -type f -name *.proto)

	go test ./...
	go mod tidy

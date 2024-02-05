tools:
	go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

proto:
	protoc --proto_path=api/ --go_out=./internal/gen --go_opt=paths=source_relative api/*.proto

weaver:
	weaver generate ./...

phony: tools

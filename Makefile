default: generate

tools:
	go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

proto:
	# Ensure that the `require_unimplemented_servers` flag is set to false for GRPC generation.
	# If left as default, additional helper methods to ensure that methods are always implemented (requireing a unimplemented server stub)
	# are added to the GRPC server interface. This means that errors from unimplemented methods are caught at runtime, not compile time.
	# It also does not allow us to generate a serviceweaver component that adheres to the server interface, as those methods are unexported.
	# Upstream context can be found here: https://github.com/grpc/grpc-go/issues/3669
	protoc --proto_path=api/ \
		--go_out=internal/gen --go_opt=paths=source_relative \
		--go-grpc_out=internal/gen --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		--grpc-gateway_out=./internal/gen --grpc-gateway_opt=paths=source_relative \
		movie.proto
	protoc --proto_path=api/ \
		--go_out=internal/gen --go_opt=paths=source_relative \
		--go-grpc_out=internal/gen --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		--grpc-gateway_out=./internal/gen --grpc-gateway_opt=paths=source_relative \
		director.proto

weaver:
	weaver generate ./...

clean:
	find . -name weaver_gen.go -exec rm {} \;
	find . -name *.pb.go -exec rm {} \;

generate: proto weaver


phony: tools proto weaver clean generate default

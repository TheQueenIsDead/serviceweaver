# serviceweaver

This repo is a demo application that is being used to test the [serviceweaver](https://serviceweaver.dev) framework 

The goal is to setup a project using the framework, that leverages as much code generation as possible once given service
proto definitions.

## Start

Retrieve all the dependencies for the project with:

```console
make tools
```

Generate the latest version of the code with:
```console
make
```

Run the app

```console
go run ./...
```

## Test

Rudimentary data can be retrieved with the example HTTP requests found in [api.http](./api.http):

## Links

 - [serviceweaver](https://serviceweaver.dev)
 - [awesome-grpc](https://github.com/grpc-ecosystem/awesome-grpc)
 - [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
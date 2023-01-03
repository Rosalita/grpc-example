# grpc-server

A simple client server grpc example.

## Generating grpc and protobuf files

Install the grpc plugin
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

compile the `chat.proto` file can use
`protoc --go_out=. --go-grpc_out=. chat.proto`

but this will compile it in the relative folder.
`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative chat.proto`

`chat.pb.go` contains the protobuf types for encoding and decoding messages.
`chat_grpc.pb.go` contains the generated rpc client and server code.

## Server

A new grpc server has no services registered to it. Have to create, then register services. 

Note that the internal server type must embed the generated grpc `UnimplementedxxxService` so that it can be passed to the generated
`RegisterxxxService` function.

Routing is defined in the `.proto` file.

## Client 

The generated grpc code contains a function to instantiate a new grpc client.

Note that the types used by the client are defined in the `.proto` file.

## Design considerations and trade offs.

The `.proto` files are the contract between the client and the server as such, they are shared by both.
To decouple client and server code from `.proto` files, its recommended to save them in a separate repository.

By keeping all the `.proto` files in one place, they can be linted in CI with something like [protoc-gen-lint](https://github.com/ckaznocha/protoc-gen-lint)
It's a good idea to version `.proto` files. The generated code from `.proto` files can also be packaged into libraries.

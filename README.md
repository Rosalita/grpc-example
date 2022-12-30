# grpc-server

A simple example of setting up a grpc server.

Install the grpc plugin
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

compile the .proto file
`protoc --go_out=. --go-grpc_out=. chat.proto`

or 
`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative chat.proto`

To set up a grpc server a new server is created then services are registered to that grpc server.

Note that the internal server type must embed the generated grpc `UnimplementedxxxService` so that it can be passed to the generated
`RegisterxxxService` function.

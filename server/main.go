package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/Rosalita/grpc-server/pb"
)

func main(){

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	s := NewServer()

	pb.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
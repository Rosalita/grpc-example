package main

import (
	"log"

	"context"
	"github.com/Rosalita/grpc-server/pb"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func NewServer() Server{
	return Server{}
}

func (s *Server) SayHello(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}
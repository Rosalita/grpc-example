package main

import (
	"log"
	"context"

	"github.com/Rosalita/grpc-server/pb"
	"google.golang.org/grpc"
)

func main(){

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &pb.Message{Body: "Hello from a client!"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.Body)
}
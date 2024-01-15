package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"log"
	"time"
)

func doLongGreet(client pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Victor"},
		{FirstName: "Maykon"},
		{FirstName: "Bruno"},
		{FirstName: "Gustavo"},
	}

	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Erro while calling Long Greet Service: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	recv, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("Long greet: %s\n", recv.Result)
}

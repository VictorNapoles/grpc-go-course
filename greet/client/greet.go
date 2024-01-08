package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"log"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	response, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Victor",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greetting: %s\n", response.Result)
}

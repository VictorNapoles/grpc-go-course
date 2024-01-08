package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(client pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")
	response, err := client.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Victor",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	for {
		res, err := response.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not greet: %v\n", err)
		}

		log.Printf("Greetting: %s\n", res.Result)
	}

}

package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"log"
)

func doAverage(client pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")

	numbers := []int64{1, 2, 3, 4, 5}
	stream, err := client.Average(context.Background())

	if err != nil {
		log.Fatalf("Erro while calling Average Service: %v\n", err)
	}

	for _, number := range numbers {
		log.Printf("Sending req: %v\n", number)
		stream.Send(&pb.AverageRequest{Number: number})
	}

	recv, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Average: %d\n", recv.Result)
}

package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"log"
)

func doSumCalculator(client pb.CalculatorServiceClient) {
	log.Println("doSumCalculator was invoked")
	response, err := client.Sum(context.Background(), &pb.SumRequest{
		Number1: 10,
		Number2: 20,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum result: %d\n", response.Result)
}

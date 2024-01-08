package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"io"
	"log"
)

func doPrimeDecomposition(client pb.CalculatorServiceClient) {
	log.Println("doPrimeDecomposition was invoked")
	responses, err := client.PrimeDecomposition(context.Background(), &pb.PrimeRequest{
		Value: 120,
	})
	if err != nil {
		log.Fatalf("Could not prime decomposite: %v\n", err)
	}

	for {
		res, err := responses.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not prime decomposite: %v\n", err)
		}

		log.Printf("Prime decomposition: %d\n", res.Result)
	}

}

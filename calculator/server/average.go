package main

import (
	pb "grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	result := int64(0)
	total := int64(0)

	size := int64(0)

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Averge : %d\n", result)
			return stream.SendAndClose(&pb.AverageResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while calculating average: %v\n", err)
		}

		log.Printf("Recived number: %d\n", recv.Number)

		size++
		total += recv.Number
		result = total / size
	}
}

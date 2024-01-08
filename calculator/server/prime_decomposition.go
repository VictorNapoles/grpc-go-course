package main

import (
	pb "grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) PrimeDecomposition(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeDecompositionServer) error {
	log.Printf("PrimeDecomposition was invoked: %v\n", in)
	var k int64 = 2
	number := in.Value
	for number > 1 {
		if number%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			number = number / k
		} else {
			k += 1
		}
	}

	return nil
}

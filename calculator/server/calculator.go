package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.Number1 + in.Number2,
	}, nil
}

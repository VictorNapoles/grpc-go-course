package main

import (
	"fmt"
	pb "grpc-go-course/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet was invoked")

	res := ""

	for {
		recv, err := stream.Recv()

		if err == io.EOF {
			stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Erro while reading response: %s/n", err)
		}

		res += fmt.Sprintf("Hello %s!/n", recv.FirstName)
	}
}

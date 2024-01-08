package main

import (
	"google.golang.org/grpc"
	pb "grpc-go-course/calculator/proto"
	"log"
	"net"
)

var addr = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalln("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: %v\n", err)
	}
}

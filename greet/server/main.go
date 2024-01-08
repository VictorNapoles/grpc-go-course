package main

import (
	"google.golang.org/grpc"
	pb "grpc-go-course/greet/proto"
	"log"
	"net"
)

var addr = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalln("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: %v\n", err)
	}

}

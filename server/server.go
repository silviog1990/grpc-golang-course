package main

import (
	"fmt"
	"log"
	"net"

	"github.com/silviog1990/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	serverIP := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", serverIP)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	fmt.Println("Server listen to:", serverIP)
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

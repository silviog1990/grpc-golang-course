package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/silviog1990/grpc-sum/sumpb"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("Execution of Sum: %v", req)
	a := req.GetSum().GetA()
	b := req.GetSum().GetB()
	result := a + b
	res := &sumpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	serverIP := "0.0.0.0:50000"
	lis, err := net.Listen("tcp", serverIP)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	fmt.Printf("Listen to %v\n", serverIP)
	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

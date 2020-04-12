package main

import (
	"fmt"
	"log"
	"net"

	"github.com/silviog1990/grpc-golang-course/streaming-server/PrimeNumberDecomposition/primenumberpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumberDecomposition(req *primenumberpb.PrimeNumberRequest, stream primenumberpb.PrimeNumberService_PrimeNumberDecompositionServer) error {
	fmt.Println("PrimeNumberDecomposition called")
	n := req.Num
	// var k int32 = 2
	k := int32(2)
	for n > 1 {
		if n%k == 0 {
			n = n / k
			res := &primenumberpb.PrimeNumberResponse{
				Result: k,
			}
			stream.Send(res)
		} else {
			k = k + 1
		}
	}
	return nil
}

func main() {
	serverIP := "0.0.0.0:50000"
	lis, err := net.Listen("tcp", serverIP)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	fmt.Printf("Listen to %v\n", serverIP)
	primenumberpb.RegisterPrimeNumberServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

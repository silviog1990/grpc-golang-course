package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/silviog1990/grpc-golang-course/bidirectional-streaming/FindMax/findmaxpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) FindMax(stream findmaxpb.FindMaxService_FindMaxServer) error {
	fmt.Println("FindMax invoked")
	numbers := []int32{}
	max := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while receiving stream: %v\n", err)
			return err
		}
		if len(numbers) == 0 {
			max = req.N
		}
		numbers = append(numbers, req.N)

		for _, num := range numbers {
			if num > max {
				max = num
			}
		}

		sendError := stream.Send(&findmaxpb.FindMaxResponse{Result: max})
		if sendError != nil {
			log.Fatalf("error while sending stream: %v\n", sendError)
			return err
		}
	}
}

func main() {
	serverIP := "0.0.0.0:50000"
	lis, err := net.Listen("tcp", serverIP)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	fmt.Printf("Listen to %v\n", serverIP)
	findmaxpb.RegisterFindMaxServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

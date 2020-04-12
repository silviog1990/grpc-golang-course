package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/silviog1990/grpc-golang-course/streaming-client/ComputeAverage/computeaveragepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) ComputeAverage(stream computeaveragepb.ComputeAverageService_ComputeAverageServer) error {
	fmt.Println("ComputeAverage invoked")
	summatory := int32(0)
	cont := int32(0)
	average := float64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&computeaveragepb.ComputeAverageResponse{
				Result: average,
			})
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		summatory += req.GetN()
		cont++
		average = float64(summatory) / float64(cont)
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
	computeaveragepb.RegisterComputeAverageServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

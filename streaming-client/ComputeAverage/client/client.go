package main

import (
	"context"
	"fmt"
	"log"

	"github.com/silviog1990/grpc-golang-course/streaming-client/ComputeAverage/computeaveragepb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to: %v", err)
	}
	defer cc.Close()
	c := computeaveragepb.NewComputeAverageServiceClient(cc)
	doClientStream(c)
}

func doClientStream(c computeaveragepb.ComputeAverageServiceClient) {
	fmt.Println("Start client stream invocation")
	numbers := []int32{3, 5, 10, 2, 1, 8}

	stream, err := c.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("error while calling ComputeAverage: %v", err)
	}

	for _, n := range numbers {
		req := &computeaveragepb.ComputeAverageRequest{
			N: n,
		}
		fmt.Printf("sending req: %v\n", req)
		stream.Send(req)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while reading resp: %v\n", err)
	}

	fmt.Printf("Average: %v\n", resp)

}

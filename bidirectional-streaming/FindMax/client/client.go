package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/silviog1990/grpc-golang-course/bidirectional-streaming/FindMax/findmaxpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to: %v", err)
	}
	defer cc.Close()
	c := findmaxpb.NewFindMaxServiceClient(cc)
	doBidirectionalStream(c)
}

func doBidirectionalStream(c findmaxpb.FindMaxServiceClient) {
	fmt.Println("Start client stream invocation")
	numbers := []int32{100, 3, 5, 1000, 2, 1, 800}

	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("error while calling FindMax: %v", err)
	}

	waitc := make(chan struct{})

	// send messages
	go func() {
		for _, n := range numbers {
			req := &findmaxpb.FindMaxRequest{
				N: n,
			}
			fmt.Printf("sending req: %v\n", req)
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	// receive messages
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving response from stream")
				break
			}
			fmt.Printf("Current Max: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc

}

package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/silviog1990/grpc-golang-course/streaming-server/PrimeNumberDecomposition/primenumberpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to: %v", err)
	}
	defer cc.Close()
	c := primenumberpb.NewPrimeNumberServiceClient(cc)
	doStream(c)
}

func doStream(c primenumberpb.PrimeNumberServiceClient) {
	fmt.Println("Start stream invocation")

	req := &primenumberpb.PrimeNumberRequest{
		Num: 120,
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		fmt.Printf("Response from PrimeNumberDecomposition: %d\n", msg.GetResult())
	}
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/silviog1990/grpc-golang-course/unary/sum/sumpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to: %v", err)
	}
	defer cc.Close()
	c := sumpb.NewSumServiceClient(cc)
	doUnary(c)
}

func doUnary(c sumpb.SumServiceClient) {
	fmt.Println("Start unary invocation")

	req := &sumpb.SumRequest{
		Sum: &sumpb.Sum{
			A: 10,
			B: 3,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	fmt.Printf("Response from Sum: %d\n", res.GetResult())
}

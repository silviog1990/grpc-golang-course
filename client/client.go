package main

import (
	"fmt"
	"log"

	"github.com/silviog1990/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to: %v", err)
	}
	defer cc.Close()
	greetpb.NewGreetServiceClient(cc)
	fmt.Println("Created client")
}

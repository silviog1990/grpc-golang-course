package main

import (
	"context"
	"log"

	"github.com/silviog1990/grpc-golang-course/blog-mongo/blogpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to: %v", err)
	}
	defer cc.Close()

	client := blogpb.NewBlogServiceClient(cc)
	client.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{})
}

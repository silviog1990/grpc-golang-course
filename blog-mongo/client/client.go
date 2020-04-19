package main

import (
	"context"
	"fmt"
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
	createBlog(client)
}

func createBlog(client blogpb.BlogServiceClient) {
	fmt.Println("create blog api invoked")
	res, err := client.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			Author:  "Silvio",
			Title:   "grpc tutorial",
			Content: "this is grpc tutorial with golang and mongodb",
		},
	})
	if err != nil {
		log.Fatalf("Error while calling CreateBlog: %v", err)
	}
	fmt.Printf("Blog created: %v\n", res)
}

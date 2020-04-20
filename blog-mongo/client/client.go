package main

import (
	"context"
	"fmt"
	"log"

	"github.com/silviog1990/grpc-golang-course/blog-mongo/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to: %v", err)
	}
	defer cc.Close()

	client := blogpb.NewBlogServiceClient(cc)
	createBlog(client)
	readBlog(client, "5e9cc96900733582dac66aed")
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

//5e9cc96900733582dac66aed
func readBlog(client blogpb.BlogServiceClient, id string) {
	fmt.Println("read blog api invoked")
	res, err := client.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		Id: id,
	})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.NotFound {
			fmt.Printf("Blog with id: %v not found\n", id)
		} else {
			log.Fatalf("Error while calling ReadBlog: %v", err)
		}
		return
	}
	fmt.Printf("Blog found: %v\n", res)
}

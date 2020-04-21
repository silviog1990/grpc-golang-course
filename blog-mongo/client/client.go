package main

import (
	"context"
	"fmt"
	"io"
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

	blogInput := &blogpb.Blog{
		Author:  "Silvio",
		Title:   "grpc tutorial",
		Content: "this is grpc tutorial with golang and mongodb",
	}

	client := blogpb.NewBlogServiceClient(cc)
	blog, _ := createBlog(client, blogInput)
	readBlog(client, blog.Id)
	updateBlog(client, blog)
	deleteBlog(client, blog.Id)
	readBlog(client, blog.Id)
	listBlogs(client)
}

func createBlog(client blogpb.BlogServiceClient, blog *blogpb.Blog) (*blogpb.Blog, error) {
	fmt.Println("create blog api invoked")
	res, err := client.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("Error while calling CreateBlog: %v", err)
		return nil, err
	}
	fmt.Printf("Blog created: %v\n", res)
	return res.Blog, nil
}

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

func updateBlog(client blogpb.BlogServiceClient, blog *blogpb.Blog) {
	fmt.Println("update blog api invoked")
	blog.Title += " MODIFIED"
	blog.Content += " MODIFIED"
	res, err := client.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("Error while calling UpdateBlog: %v", err)
	}
	fmt.Printf("Blog created: %v\n", res)
}

func deleteBlog(client blogpb.BlogServiceClient, id string) {
	fmt.Println("delete blog api invoked")
	res, err := client.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
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
	fmt.Printf("Blog deleted: %v\n", res)
}

func listBlogs(client blogpb.BlogServiceClient) {
	fmt.Println("list blogs api invoked")
	stream, err := client.ListBlogs(context.Background(), &blogpb.ListBlogsRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlogs RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}

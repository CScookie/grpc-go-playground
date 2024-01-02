package main

import (
	"context"
	"log"

	pb "github.com/cscookie/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog() was invoked")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error while reading: %v\n", err)
	}
	log.Printf("Blog read: %s\n", res)
	return res
}

package main

import (
	"context"
	"log"

	pb "github.com/cscookie/grpc-go-course/blog/proto"
)

func DeleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("DeleteBlog() was invoked")

	req := &pb.BlogId{Id: id}
	_, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error while deleting: %v\n", err)
	}
	log.Printf("Blog deleted")
}

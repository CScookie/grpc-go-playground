package main

import (
	"context"
	"log"

	pb "github.com/cscookie/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("readBlog() was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Jeremy",
		Title:    "A new title",
		Content:  "Updated content",
	}
	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error while update: %v\n", err)
	}
	log.Println("Blog was updated")

}

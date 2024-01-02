package main

import (
	"context"
	"log"
	"time"

	pb "github.com/cscookie/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline() was invokved")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "Jeremy",
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %s\n", e.Message())
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("doGreetWithDeadline() Greeting: %s\n", res.Result)
}

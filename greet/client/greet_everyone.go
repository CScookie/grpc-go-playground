package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/cscookie/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Printf("doGreetEveryone() was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creaitng stream: %v\n", err)
	}
	reqs := []*pb.GreetRequest{
		{FirstName: "Jeremy"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}
	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}

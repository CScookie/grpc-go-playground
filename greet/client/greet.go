package main

import (
	"context"
	"log"

	pb "github.com/cscookie/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invokved")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Jeremy",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}

package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cscookie/grpc-go-course/sum/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %\n", err)
	}
	defer conn.Close()

	c := pb.NewSumServiceClient(conn)

	callSumService(c)
}

func callSumService(c pb.SumServiceClient) {
	log.Println("callSumService was invokved")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstInt:  12,
		SecondInt: 11,
	})
	if err != nil {
		log.Fatalf("Could not Sum: %v\n", err)
	}
	log.Printf("Sum result: %v\n", res.Result)
}

package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pb "github.com/cscookie/grpc-go-course/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	callSqrt(c)
}

func callSqrt(c pb.CalculatorServiceClient) {
	log.Println("callSqrt() was invokved")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: -2,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("A negative number was sent to server")
			}
			return
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
		log.Fatalf("Could not Sqrt: %v\n", err)
	}
	log.Printf("Sqrt() result: %v\n", res.Result)
}

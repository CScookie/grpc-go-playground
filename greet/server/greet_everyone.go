package main

import (
	"io"
	"log"

	pb "github.com/cscookie/grpc-go-course/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone() was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client steam: %v\n", err)
		}
		res := "Hello" + req.FirstName + "!"
		if err := stream.Send(&pb.GreetResponse{
			Result: res,
		}); err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}

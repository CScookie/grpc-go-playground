package main

import (
	"fmt"
	"log"

	pb "github.com/cscookie/grpc-go-course/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTies function was invoked with :%v\n", in)

	for i := 0; i < 100; i++ {
		res := fmt.Sprintf("Hello %s, number %v", in.GetFirstName(), i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}

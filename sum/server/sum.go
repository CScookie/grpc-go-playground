package main

import (
	"context"
	"log"

	pb "github.com/cscookie/grpc-go-course/sum/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet fuction was invoked with %v\n", in)
	var result_int uint32 = in.GetFirstInt() + in.GetSecondInt()

	log.Printf("Sum of %v and %v: %v\n", in.GetFirstInt(), in.GetSecondInt(), result_int)

	return &pb.SumResponse{
		Result: result_int,
	}, nil
}

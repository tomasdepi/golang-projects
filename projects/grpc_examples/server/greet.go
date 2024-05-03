package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/tomasdepi/golang/projects/grpc_examples/pb/greet"
)

func (s *MyGreetServiceServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked by %s\n", req.Name)
	return &pb.GreetResponse{
		Result: "Hello " + req.Name,
	}, nil
}

func (s *MyGreetServiceServer) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked\n")
	var result string
	for i := 0; i < 10; i++ {
		result = fmt.Sprintf("Hello %s number %d", req.Name, i)
		stream.Send(&pb.GreetResponse{
			Result: result,
		})
	}

	return nil
}

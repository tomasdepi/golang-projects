package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

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

func (s *MyGreetServiceServer) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked\n")

	names := []string{}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			stream.SendAndClose(&pb.GreetResponse{
				Result: "Hello to " + strings.Join(names, ", "),
			})
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %s\n", err)
		}

		names = append(names, msg.Name)
	}

	return nil
}

func (s *MyGreetServiceServer) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked\n")

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %s\n", err)
		}

		res := msg.Name

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending response %s\n", err)
		}
	}

	return nil
}

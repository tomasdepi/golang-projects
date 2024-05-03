package main

import (
	"context"
	"io"
	"log"

	pb "github.com/tomasdepi/golang/projects/grpc_examples/pb/greet"
)

func doGreet(gsc pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := gsc.Greet(context.Background(), &pb.GreetRequest{
		Name: "Depi",
	})

	if err != nil {
		log.Fatalf("Failed to Greet %s\n", err)
	}

	log.Printf("Greet %s\n", res)
}

func doGreetManyTimes(gsc pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	greetRequest := &pb.GreetRequest{
		Name: "Depi",
	}

	stream, err := gsc.GreetManyTimes(context.Background(), greetRequest)

	if err != nil {
		log.Fatalf("Error calling GreetManyTimes %s\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return
		}

		if err != nil {
			log.Fatalf("Error while reading stream %s\n", err)
		}

		log.Printf("%s", msg.Result)
	}

}

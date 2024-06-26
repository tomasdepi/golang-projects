package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/tomasdepi/golang/projects/grpc_examples/pb/greet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func doLongGreet(gsc pb.GreetServiceClient, names []string) {
	log.Println("doLongGreet was invoked")

	stream, err := gsc.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error calling GreetManyTimes %s\n", err)
	}

	for _, value := range names {
		stream.Send(&pb.GreetRequest{
			Name: value,
		})
	}

	response, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while reading response %s\n", err)
	}

	log.Printf("Long Greet: %s\n", response.Result)
}

func doGreetEveryone(gsc pb.GreetServiceClient, names []string) {
	log.Println("doGreetEveryone was invoked")

	stream, err := gsc.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error calling GreetEveryone %s\n", err)
	}

	waitc := make(chan []string)

	go func() {
		for _, name := range names {
			stream.Send(&pb.GreetRequest{
				Name: name,
			})
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
				log.Fatalf("Error while reading stream %s\n", err)
				break
			}

			log.Printf("Received: %s\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}

func doGreetWithDeadline(gsc pb.GreetServiceClient, name string, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		Name: name,
	}

	res, err := gsc.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)
		if e.Code() == codes.DeadlineExceeded {
			log.Println("Deadline Exceeded")
			return
		} else {
			log.Panicf("Unexpected gRPC error: %v\n", err)
		}
		if ok {

		} else {
			log.Fatalf("A non gRCP error: %v\n", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}

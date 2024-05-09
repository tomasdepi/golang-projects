package main

import (
	"log"
	"net"

	calc "github.com/tomasdepi/golang/projects/grpc_examples/pb/calculator"
	pb "github.com/tomasdepi/golang/projects/grpc_examples/pb/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50001"

type MyGreetServiceServer struct {
	pb.GreetServiceServer
}

type MyCalculatorServiceServer struct {
	calc.CalculatorServiceServer
}

func main() {

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Unable to Listen %s\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	tls := true
	opts := []grpc.ServerOption{}

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed to load certificates %s\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	greetService := &MyGreetServiceServer{}
	calcService := &MyCalculatorServiceServer{}

	pb.RegisterGreetServiceServer(s, greetService)
	calc.RegisterCalculatorServiceServer(s, calcService)

	err = s.Serve(listener)

	if err != nil {
		log.Fatalf("Unable to Serve %s\n", err)
	}

}

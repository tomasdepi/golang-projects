package main

import (
	"log"

	calc "github.com/tomasdepi/golang/projects/grpc_examples/pb/calculator"
	pb "github.com/tomasdepi/golang/projects/grpc_examples/pb/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverEndpoint string = "localhost:50001"

func main() {
	conn, err := grpc.Dial(serverEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to Connect %s\n", err)
	}

	defer conn.Close()

	clientgreet := pb.NewGreetServiceClient(conn)
	clientcalc := calc.NewCalculatorServiceClient(conn)

	doGreet(clientgreet) // Unary call

	res := doSum(clientcalc, 10, 3) // Unary call
	log.Printf("The result is %v\n", res)
	res = doMul(clientcalc, 7, 4) // Unary call
	log.Printf("The result is %v\n", res)

	doGreetManyTimes(clientgreet) // Streaming Server call

	doPrime(clientcalc, 120)
}

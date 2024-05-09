package main

import (
	"log"
	"time"

	calc "github.com/tomasdepi/golang/projects/grpc_examples/pb/calculator"
	pb "github.com/tomasdepi/golang/projects/grpc_examples/pb/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var serverEndpoint string = "localhost:50001"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		// keyFile := "ssl/server.key"

		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Failed to load creds: %s\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(serverEndpoint, opts...)

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

	doPrime(clientcalc, 120) // Streaming Server call

	doLongGreet(clientgreet, []string{"Okarin", "Calabacita", "Flaco"}) // Client Server call

	doAvg(clientcalc, []uint64{1, 2, 3, 4}) // Client Server call

	doGreetEveryone(clientgreet, []string{"Okarin", "Calabacita", "Flaco"}) // Bi-lateral Streaming

	doMax(clientcalc, []uint64{1, 5, 3, 6, 2, 20}) // Bi-lateral Streaming

	doSqrt(clientcalc, 10) // Error Handling

	doSqrt(clientcalc, -10) // Error Handling

	doGreetWithDeadline(clientgreet, "Depi", 5*time.Second) // Timeout Example
}

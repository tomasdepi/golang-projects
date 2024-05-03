package main

import (
	"context"
	"io"
	"log"

	calc "github.com/tomasdepi/golang/projects/grpc_examples/pb/calculator"
)

func doSum(ssc calc.CalculatorServiceClient, a int64, b int64) int64 {
	log.Println("doSum was invoked")
	res, err := ssc.Sum(context.Background(), &calc.SumRequest{
		A: a,
		B: b,
	})

	if err != nil {
		log.Fatalf("Failed to Sum %s\n", err)
	}

	return res.Result
}

func doMul(ssc calc.CalculatorServiceClient, a int64, b int64) int64 {
	log.Println("doMul was invoked")
	res, err := ssc.Mul(context.Background(), &calc.MulRequest{
		A: a,
		B: b,
	})

	if err != nil {
		log.Fatalf("Failed to Mul %s\n", err)
	}

	return res.Result
}

func doPrime(ssc calc.CalculatorServiceClient, number int64) {
	log.Println("doPrime was invoked")

	req := calc.PrimeRequest{
		Number: number,
	}

	stream, err := ssc.Prime(context.Background(), &req)

	if err != nil {
		log.Fatalf("Failed to Prime %s\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return
		}

		if err != nil {
			log.Fatalf("Error while reading stream %s\n", err)
		}

		log.Printf("We got the prime %d", msg.Prime)
	}
}

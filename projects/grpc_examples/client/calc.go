package main

import (
	"context"
	"io"
	"log"
	"time"

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

func doAvg(ssc calc.CalculatorServiceClient, numbers []uint64) {
	log.Println("doAvg was invoked")

	stream, err := ssc.Avg(context.Background())

	if err != nil {
		log.Fatalf("Failed to Avg %s\n", err)
	}

	for _, n := range numbers {
		stream.Send(&calc.AvgRequest{
			Number: n,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while reading response from Avg %s\n", err)
	}

	log.Printf("We got the avg %f", res.Response)
}

func doMax(ssc calc.CalculatorServiceClient, numbers []uint64) {
	log.Println("doMax was invoked")

	stream, err := ssc.Max(context.Background())

	if err != nil {
		log.Fatalf("Error calling Max %s\n", err)
	}

	waitc := make(chan []uint64)

	go func() {
		for _, number := range numbers {
			stream.Send(&calc.MaxRequest{
				Number: number,
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

			log.Printf("Received New Max Number: %d\n", res.Response)
		}

		close(waitc)
	}()

	<-waitc
}

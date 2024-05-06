package main

import (
	"context"
	"io"
	"log"

	calc "github.com/tomasdepi/golang/projects/grpc_examples/pb/calculator"
)

func getPrimes(n int64) []int64 {
	primes := []int64{}
	var k int64 = 2
	for n > 1 {
		if n%k == 0 {
			primes = append(primes, k)
			n = n / k
		} else {
			k = k + 1
		}
	}

	return primes
}

func calculateAverage(list []uint64) float64 {
	var total uint64 = 0
	for _, num := range list {
		total += num
	}
	return float64(total) / float64(len(list))
}

func (c *MyCalculatorServiceServer) Sum(ctx context.Context, req *calc.SumRequest) (*calc.SumResponse, error) {
	log.Printf("Sum function was invoked")
	return &calc.SumResponse{
		Result: req.A + req.B,
	}, nil
}

func (c *MyCalculatorServiceServer) Mul(ctx context.Context, req *calc.MulRequest) (*calc.MulResponse, error) {
	log.Printf("Mul function was invoked")
	return &calc.MulResponse{
		Result: req.A * req.B,
	}, nil
}

func (c *MyCalculatorServiceServer) Prime(req *calc.PrimeRequest, stream calc.CalculatorService_PrimeServer) error {
	log.Printf("Prime function was invoked")

	primes := getPrimes(req.Number)

	for _, value := range primes {
		stream.Send(&calc.PrimeResponse{
			Prime: value,
		})
	}

	return nil
}

func (c *MyCalculatorServiceServer) Avg(stream calc.CalculatorService_AvgServer) error {
	log.Printf("Avg function was invoked")

	numbers := []uint64{}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			stream.SendAndClose(&calc.AvgResponse{
				Response: calculateAverage(numbers),
			})
			break
		}

		if err != nil {
			log.Fatalf("Failed to receive from Avg %s\n", err)
		}

		numbers = append(numbers, msg.Number)
	}

	return nil
}

package main

import (
	"context"
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

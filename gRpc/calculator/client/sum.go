package main

import (
	"context"
	"log"

	pb "github.com/luisfelipegodoi/study/gRpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}

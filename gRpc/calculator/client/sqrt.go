package main

import (
	"context"
	"log"

	pb "github.com/luisfelipegodoi/study/gRpc/calculator/proto"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})
	if err != nil {

	}
}

package main

import (
	"context"
	"log"
	"time"

	pb "github.com/luisfelipegodoi/study/gRpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Luixx",
	}

	res, err := c.GreetWithDeadLine(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
				return
			} else {
				log.Fatalf("Unexpeted gRpc error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRpc error: %v\n", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}

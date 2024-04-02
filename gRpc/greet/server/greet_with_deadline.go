package main

import (
	"context"
	"log"

	pb "github.com/luisfelipegodoi/study/gRpc/greet/proto"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", in)
}

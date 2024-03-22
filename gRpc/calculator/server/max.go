package main

import (
	"io"
	"log"

	pb "github.com/luisfelipegodoi/study/gRpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("")
	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("")
			break
		}
	}
}

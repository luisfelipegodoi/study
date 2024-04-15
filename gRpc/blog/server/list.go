package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/luisfelipegodoi/study/gRpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("ListBlogs was invoked")

	ctx := context.Background()

	cur, err := collection.Find(ctx, primitive.D{})
	if err != nil{
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx){
		data := &BlogItem{}
		err := 
	}
}

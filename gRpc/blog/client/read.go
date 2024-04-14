package client

import (
	"log"

	pb "github.com/luisfelipegodoi/study/gRpc/blog/proto"
	"golang.org/x/net/context"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("--- readBlog was invoked ---")

	ctx := context.Background()
	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(ctx, req)
	if err != nil {
		log.Printf("Error happened while reading: %v \n", err)
	}

	log.Printf("Blog was read: %v \n", res)

	return res
}

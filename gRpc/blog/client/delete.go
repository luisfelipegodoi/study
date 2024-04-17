package client

import (
	"context"
	"log"

	pb "github.com/luisfelipegodoi/study/gRpc/blog/proto"
)

func deleteBloog(c pb.BlogServiceClient, id string) {
	log.Println("--- deleteBlog was invoked ---")

	ctx := context.Background()

	_, err := c.DeleteBlog(ctx, &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}

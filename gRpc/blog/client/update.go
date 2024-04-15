package client

import (
	"context"
	"log"

	pb "github.com/luisfelipegodoi/study/gRpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("--- updateBlog was invoked ---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Luix",
		Title:    "New Title of blog",
		Content:  "Content of the first blog",
	}

	ctx := context.Background()

	_, err := c.UpdateBlog(ctx, newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}

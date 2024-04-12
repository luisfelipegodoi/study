package client

import (
	"context"
	"log"

	pb "github.com/luisfelipegodoi/study/gRpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--- createBlog was invoked ---")

	blog := &pb.Blog{
		AuthorId: "Luix",
		Title:    "Title of Blog",
		Content:  "Content of Blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error %v\n", err)
	}

	log.Println("Blog has been created: %s\n", res.Id)

	return res.Id
}

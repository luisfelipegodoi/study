package client

import (
	"log"

	pb "github.com/luisfelipegodoi/study/gRpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("--- updateBlog was invoked ---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Luix",
		Title:    "New Title of blog",
	}
}

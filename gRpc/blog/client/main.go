package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/luisfelipegodoi/study/gRpc/blog/proto"
)

var addr = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)            // id valid
	readBlog(c, "aNonValidId") // invalid id
	updateBlog(c, id)
}

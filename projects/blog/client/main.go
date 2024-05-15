package main

import (
	"log"

	pbblog "github.com/tomasdepi/golang/projects/blog/pbblog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverEndpoint string = "localhost:50001"

func main() {

	conn, err := grpc.Dial(serverEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to Connect %s\n", err)
	}

	defer conn.Close()

	clientblog := pbblog.NewBlogServiceClient(conn)

	id := createBlog(clientblog)
	blog := readBlog(clientblog, id) // valid
	log.Printf("Blog was read, the author is %s\n", blog.AuthorId)
	updateBlog(clientblog, id)
	listBlog(clientblog)
	deleteBlog(clientblog, id)
	// readBlog(clientblog, "gf87sdg8s7yg89") // random value
}

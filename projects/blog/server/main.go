package main

import (
	"context"
	"log"
	"net"

	pb "github.com/tomasdepi/golang/projects/blog/pbblog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50001"

type MyBlogServiceServer struct {
	pb.BlogServiceServer
}

func main() {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@127.0.0.1:27017/"))

	if err != nil {
		log.Fatalf("Unable to initialize mongo client %s", err)
	}

	err = mongoClient.Connect(context.Background())

	if err != nil {
		log.Fatalf("Unable to connect to mongo %s", err)
	}

	collection = mongoClient.Database("blogdb").Collection("blog")

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Unable to Listen %s\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()

	blogService := &MyBlogServiceServer{}

	pb.RegisterBlogServiceServer(s, blogService)

	err = s.Serve(listener)

	if err != nil {
		log.Fatalf("Unable to Serve %s\n", err)
	}

}

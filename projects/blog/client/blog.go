package main

import (
	"context"
	"io"
	"log"

	pbblog "github.com/tomasdepi/golang/projects/blog/pbblog"
	"google.golang.org/protobuf/types/known/emptypb"
)

func createBlog(c pbblog.BlogServiceClient) string {
	blog := &pbblog.Blog{
		AuthorId: "Depi",
		Title:    "Title",
		Contect:  "My First Blog",
	}

	blogResponse, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Failed creating the blog %v\n", err)
	}

	log.Printf("Created blog: %s\n", blogResponse.Id)

	return blogResponse.Id
}

func readBlog(c pbblog.BlogServiceClient, id string) *pbblog.Blog {

	req := &pbblog.BlogId{Id: id}
	blogResponse, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while reading response: %v\n", err)
	}

	return blogResponse
}

func updateBlog(c pbblog.BlogServiceClient, id string) {
	newBlog := &pbblog.Blog{
		Id:       id,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Contect:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}

func listBlog(c pbblog.BlogServiceClient) {
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}

func deleteBlog(c pbblog.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")
	_, err := c.DeleteBlog(context.Background(), &pbblog.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}

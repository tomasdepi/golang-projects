package main

import (
	"context"
	"fmt"
	"log"

	pbblog "github.com/tomasdepi/golang/projects/blog/pbblog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *MyBlogServiceServer) CreateBlog(ctx context.Context, blog *pbblog.Blog) (*pbblog.BlogId, error) {
	log.Printf("Create Blog was invoked with %v\n", blog)

	data := &BlogItem{
		AuthorID: blog.AuthorId,
		Content:  blog.Contect,
		Title:    blog.Title,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unable to insert blog %v\n", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot conver to OID"))
	}

	return &pbblog.BlogId{
		Id: oid.Hex(),
	}, nil
}

func (s *MyBlogServiceServer) ReadBlog(ctx context.Context, blog *pbblog.BlogId) (*pbblog.Blog, error) {
	oid, err := primitive.ObjectIDFromHex(blog.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprint("Unable to parse Id %v\n", err))
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	err = res.Decode(data)

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound, "Cannot find blog with provided ID",
		)
	}

	return documentToBlog(data), nil
}

func (*MyBlogServiceServer) UpdateBlog(ctx context.Context, in *pbblog.Blog) (*emptypb.Empty, error) {

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Contect,
	}
	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data}, // all elements we want to update
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with ID",
		)
	}

	return &emptypb.Empty{}, nil
}

func (*MyBlogServiceServer) ListBlogs(_ *emptypb.Empty, stream pbblog.BlogService_ListBlogsServer) error {

	ctx := context.Background()
	cur, err := collection.Find(ctx, primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil
}

func (*MyBlogServiceServer) DeleteBlog(ctx context.Context, in *pbblog.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog was not found",
		)
	}

	return &emptypb.Empty{}, nil
}

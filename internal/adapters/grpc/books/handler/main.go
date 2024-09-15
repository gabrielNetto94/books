package handler

import (
	"context"
	"fmt"

	pb "books/internal/adapters/grpc/books/proto"
)

func GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {

	fmt.Printf("req.Id: %v\n", req.Id)

	return &pb.GetBookResponse{
		Id:     "23",
		Title:  "Test Title",
		Author: "Test Author",
		Desc:   "Test Description",
	}, nil

}

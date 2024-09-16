package handler

import (
	"context"

	pb "books/internal/adapters/grpc/books/proto"
	"books/internal/core/services"
)

type Server struct {
	pb.BookServiceServer
	Service *services.BookService
}

func (s Server) GetBook(_ context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {

	book, serviceErr := s.Service.FindById(req.Id)
	if serviceErr.Error != nil {
		return &pb.GetBookResponse{}, serviceErr.Error
	}

	return &pb.GetBookResponse{
		Id:     req.Id,
		Title:  book.Title,
		Author: book.Author,
		Desc:   book.Desc,
	}, nil
}

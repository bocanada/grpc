package server

import (
	"context"

	"github.com/bocanada/grpc/models"
	"github.com/bocanada/grpc/repository"
	"github.com/bocanada/grpc/studentpb"
)

type Server struct {
	repo repository.Respository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Respository) *Server {
	return &Server{repo, studentpb.UnimplementedStudentServiceServer{}}
}

func (s *Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Server) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	if err := s.repo.SetStudent(ctx, &models.Student{Id: req.Id, Name: req.Name, Age: req.Age}); err != nil {
		return nil, err
	}
	return &studentpb.SetStudentResponse{Id: req.Id}, nil
}

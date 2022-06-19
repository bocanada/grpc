package server

import (
	"context"
	"errors"
	"io"

	"github.com/bocanada/grpc/models"
	"github.com/bocanada/grpc/repository"
	"github.com/bocanada/grpc/testpb"
)

type TestServer struct {
	repo repository.Respository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Respository) *TestServer {
	return &TestServer{repo, testpb.UnimplementedTestServiceServer{}}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	if err := s.repo.SetTest(ctx, &models.Test{Id: req.Id, Name: req.Name}); err != nil {
		return nil, err
	}
	return &testpb.SetTestResponse{Id: req.Id}, nil
}

func (s *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return stream.SendAndClose(&testpb.SetQuestionResponse{Ok: true})
		}
		if err != nil {
			return err
		}
		qa := &models.Question{Id: msg.Id, Question: msg.Question, Answer: msg.Answer, TestId: msg.TestId}
		if err := s.repo.SetQuestion(stream.Context(), qa); err != nil {
			return err
		}
	}
}

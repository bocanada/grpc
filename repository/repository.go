package repository

import (
	"context"

	"github.com/bocanada/grpc/models"
)

type Respository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

var impl Respository

func SetRepository(repo Respository) {
	impl = repo
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return impl.GetStudent(ctx, id)
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return impl.SetStudent(ctx, student)
}

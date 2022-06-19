package repository

import (
	"context"

	"github.com/bocanada/grpc/models"
)

type Respository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) error
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

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return impl.GetTest(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return impl.SetTest(ctx, test)
}

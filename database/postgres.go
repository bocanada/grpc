package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/bocanada/grpc/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (r *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	return err
}

func (r *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var student models.Student
	for rows.Next() {
		if err := rows.Scan(&student.Id, &student.Name, &student.Age); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.Id, test.Name)
	return err
}

func (r *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var test models.Test
	for rows.Next() {
		if err := rows.Scan(&test.Id, &test.Name); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &test, nil
}

func (r *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO questions (id, question, answer, test_id) VALUES ($1, $2, $3, $4)", question.Id, question.Question, question.Answer, question.TestId)
	return err
}

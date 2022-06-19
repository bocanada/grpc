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

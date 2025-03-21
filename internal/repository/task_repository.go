package repository

import (
	"context"
	"errors"
	"skillsrock-todo/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	GetAll(ctx context.Context) ([]model.Task, error)
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id int64) error
}

type taskRepo struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(ctx context.Context, task *model.Task) error {
	query := `INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id, status, created_at, updated_at`
	err := r.db.QueryRow(ctx, query, task.Title, task.Description).Scan(&task.ID, &task.Status, &task.CreatedAt, &task.UpdatedAt)

	return err
}

func (r *taskRepo) GetAll(ctx context.Context) ([]model.Task, error) {
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskRepo) Update(ctx context.Context, task *model.Task) error {
	query := `UPDATE tasks SET title = $1, description = $2, status = $3, updated_at=NOW() WHERE id=$4`
	commandTag, err := r.db.Exec(ctx, query, task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("[INFO] Task not found")
	}

	return nil
}

func (r *taskRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM tasks WHERE id=$1`
	commandTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("[INFO] Task not found")
	}
	return nil
}

package service

import (
	"context"
	"errors"
	"skillsrock-todo/internal/model"
	"skillsrock-todo/internal/repository"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *model.Task) error
	GetTasks(ctx context.Context) ([]model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, id int64) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context, task *model.Task) error {
	if task.Title == "" {
		return errors.New("[ERROR] Title is required")
	}

	return s.repo.Create(ctx, task)
}

func (s *taskService) GetTasks(ctx context.Context) ([]model.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *taskService) UpdateTask(ctx context.Context, task *model.Task) error {
	if task.ID == 0 {
		return errors.New("[ERROR] ID is required")
	}
	return s.repo.Update(ctx, task)
}

func (s *taskService) DeleteTask(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

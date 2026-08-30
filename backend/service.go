package main

import (
	"context"
	"errors"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidTitle = errors.New("invalid title")
	ErrInvalidStatus = errors.New("invalid status")
	ErrTodoNotFound = errors.New("todo not found")
)


type TodoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) GetTodos (
	ctx context.Context,
) ([]Todo, error) {
	todos, err := s.repo.FindAll(ctx)
	if err != nil {
		return []Todo{}, err
	}

	return todos, nil
}

func (s *TodoService) GetTodo (
	ctx context.Context,
	id int,
) (Todo, error) {
	todo, found, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return Todo{}, err
	}

	if !found {
		return Todo{}, ErrTodoNotFound
	}

	return todo, nil
}

func (s *TodoService) CreateTodo (
	ctx context.Context,
	title string,
) (Todo, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Todo{}, ErrInvalidTitle
	}

	if utf8.RuneCountInString(title) > 100 {
		return Todo{}, ErrInvalidTitle
	}

	newTodo := Todo{
		Title: title,
		Status: TodoStatusTodo,
	}

	todo, err := s.repo.Create(ctx, newTodo)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) UpdateTodo (
	ctx context.Context,
	id int,
	status TodoStatus,
) (Todo, error) {
	if !status.Valid() {
		return Todo{}, ErrInvalidStatus
	}

	todo, found, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return Todo{}, err
	}

	if !found {
		return Todo{}, ErrTodoNotFound
	}

	todo.Status = status

	updatedTodo, found, err := s.repo.Update(ctx, todo)
	if err != nil {
		return Todo{}, err
	}

	if !found {
		return Todo{}, ErrTodoNotFound
	}

	return updatedTodo, nil
}

func (s *TodoService) DeleteTodo (
	ctx context.Context,
	id int,
) error {
	deleted, err := s.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	if !deleted{
		return ErrTodoNotFound
	}

	return nil
}
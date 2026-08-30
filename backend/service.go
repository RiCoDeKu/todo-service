package main

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/redis/go-redis/v9"
)

var (
	ErrInvalidTitle = errors.New("invalid title")
	ErrInvalidStatus = errors.New("invalid status")
	ErrTodoNotFound = errors.New("todo not found")
)


type TodoService struct {
	repo TodoRepository
	redis *redis.Client
}

func NewTodoService(repo TodoRepository, redisClient *redis.Client) *TodoService {
	return &TodoService{
		repo: repo,
		redis: redisClient,
	}
}

const todosCacheKey = "todos:all"

func (s *TodoService) GetTodos (
	ctx context.Context,
) ([]Todo, error) {
	cached, err := s.redis.Get(ctx, todosCacheKey).Result()

	if err == nil {
		var todos []Todo

		// Unmarshal: JSON → Go
		if err := json.Unmarshal([]byte(cached), &todos); err != nil {
			return nil, err
		}

		return todos, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	todos, err := s.repo.FindAll(ctx)
	if err != nil {
		return []Todo{}, err
	}

	// Marshal: Go → JSON
	data, err := json.Marshal(todos)
	if err != nil {
		return nil, err
	}

	if err := s.redis.Set(
		ctx,
		todosCacheKey,
		data,
		60*time.Second,
	).Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) invalidateTodosCache(
	ctx context.Context,
) error {
	return s.redis.Del(ctx, todosCacheKey).Err()
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

	if err := s.invalidateTodosCache(ctx,); err != nil {
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

	if err := s.invalidateTodosCache(ctx); err != nil {
		return Todo{}, err
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

	if err := s.invalidateTodosCache(ctx); err != nil {
		return err
	}

	return nil
}
package main

import (
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/labstack/echo/v5"
)

type Handler struct{}

var _ ServerInterface = (*Handler)(nil)

func (h *Handler) GetTodos(c *echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func (h *Handler) GetTodo(c *echo.Context, todoId int,) error {
	for _, todo := range todos {
		if todo.Id == todoId {
			return c.JSON(http.StatusOK, todo)
		}
	}

	return c.JSON(http.StatusNotFound, ErrorResponse{Message: "todo not found"})
}

func (h *Handler) CreateTodo(c *echo.Context) error {
	var req CreateTodoRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request body"})
	}

	title := strings.TrimSpace(req.Title)
	if title == "" {
		return c.JSON(
			http.StatusBadRequest, 
			ErrorResponse{
				Message: "title is required",
			},
		)
	}

	if utf8.RuneCountInString(title) > 100 {
		return c.JSON(
			http.StatusBadRequest,
			ErrorResponse{
				Message: "title must be 100 characters or less",
			},
		)
	}

	newTodo := Todo{
		Id: nextTodoId,
		Title: title,
		Status: TodoStatusTodo,
	}

	nextTodoId++

	todos = append(todos, newTodo)

	return c.JSON(http.StatusCreated, newTodo)
}

func (h *Handler) UpdateTodo(c *echo.Context, todoId int,) error {
	var req UpdateTodoRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, 
			ErrorResponse{Message: "invalid request body"},
		)
	}

	if !req.Status.Valid() {
		return c.JSON(
			http.StatusBadRequest,
			ErrorResponse{
				Message: "status must be todo or done",
			},
		)
	}

	for idx, todo := range(todos) {
		if todo.Id == todoId {
			todos[idx].Status = TodoStatus(req.Status)

			return c.JSON(http.StatusOK, todos[idx])
		}
	}

	return c.JSON(http.StatusNotFound, 
		ErrorResponse{Message: "todo not found"},
	)
}

func (h *Handler) DeleteTodo(c *echo.Context, todoId int,) error {
	for idx, todo := range todos {
		if todo.Id == todoId {
			todos = append(todos[:idx], todos[idx+1:]...)

			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, 
		ErrorResponse{Message: "todo not found"},
	)
}
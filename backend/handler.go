package main

import (
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/labstack/echo/v5"
)

func getTodos(c *echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func getTodo(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid todo id"})
	}

	for _, todo := range todos {
		if todo.ID == id {
			return c.JSON(http.StatusOK, todo)
		}
	}

	return c.JSON(http.StatusNotFound, ErrorResponse{Message: "todo not found"})
}

func createTodo(c *echo.Context) error {
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
		ID: nextTodoId,
		Title: title,
		Status: "todo",
	}

	nextTodoId++

	todos = append(todos, newTodo)

	return c.JSON(http.StatusCreated, newTodo)
}

func updateTodo(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, 
			ErrorResponse{Message: "invalid todo id"},
		)
	}

	var req UpdateTodoRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, 
			ErrorResponse{Message: "invalid request body"},
		)
	}

	if req.Status != "todo" && req.Status != "done" {
		return c.JSON(
			http.StatusBadRequest,
			ErrorResponse{
				Message: "status must be todo or done",
			},
		)
	}

	for idx, todo := range(todos) {
		if todo.ID == id {
			todos[idx].Status = req.Status

			return c.JSON(http.StatusOK, todos[idx])
		}
	}

	return c.JSON(http.StatusNotFound, 
		ErrorResponse{Message: "todo not found"},
	)
}

func deleteTodo(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, 
			ErrorResponse{Message: "invalid todo id"},
		)
	}

	for idx, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:idx], todos[idx+1:]...)

			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, 
		ErrorResponse{Message: "todo not found"},
	)
}
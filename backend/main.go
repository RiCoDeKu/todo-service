package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

// Type of Response / Internal Data
type Todo struct {
	ID		int		`json:"id"`
	Title	string	`json:"title"`
	Status	string	`json:"status"`
}

// for POST
type CreateTodoRequest struct {
	Title	string	`json:"title"`
	Status	string	`json:"status"`
}

// for PATCH
type UpdateTodoRequest struct {
	Status	string	`json:"status"`
}

func main() {
	e := echo.New()

	todos := []Todo{
	{
		ID:	1,
		Title:  "Reactを勉強する",
		Status: "done",
	},
	{
		ID:     2,
		Title:  "Goを勉強する",
		Status: "todo",
	},
}

	// GET: get todos
	e.GET("/todos", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, todos)
	})

	// GET: get todo from specified todo id
	e.GET("/todos/:id", func(c *echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid todo id")
		}

		for _, todo := range todos {
			if todo.ID == id {
				return c.JSON(http.StatusOK, todo)
			}
		}

		return c.String(http.StatusNotFound, "todo not found")
	})

	//POST: post new todo
	e.POST("/todos", func(c *echo.Context) error {
		var req CreateTodoRequest

		if err := c.Bind(&req); err != nil {
			return c.String(http.StatusBadRequest, "invalid request body")
		}

		newTodo := Todo{
			ID: len(todos) + 1,
			Title: req.Title,
			Status: req.Status,
		}

		todos = append(todos, newTodo)

		return c.JSON(http.StatusOK, newTodo)
	})
	
	//PATCH: update todo status
	e.PATCH("/todos/:id", func(c *echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "invalid todo id")
		}

		var req UpdateTodoRequest

		if err := c.Bind(&req); err != nil {
			c.String(http.StatusBadRequest, "invalid request body")
		}

		for idx, todo := range(todos) {
			if todo.ID == id {
				todos[idx].Status = req.Status

				return c.JSON(http.StatusOK, todos[idx])
			}
		}

		return c.String(http.StatusNotFound, "todo not found")
	})

	//DELETE: delete specified todo
	e.DELETE("/todos/:id", func(c *echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "invalid todo id")
		}

		for idx, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:idx], todos[idx+1:]...)

				return c.NoContent(http.StatusNoContent)
			}
		}

		return c.String(http.StatusNotFound, "todo not found")
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
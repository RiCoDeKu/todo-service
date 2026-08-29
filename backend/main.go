package main

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

var todos = []Todo{
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

var nextTodoId = 3

func main() {
	e := echo.New()

	// CORS Middleware
	e.Use(middleware.CORS("http://localhost:5173"))
	
	e.GET("/todos", getTodos)			// GET: get todos
	e.GET("/todos/:id", getTodo )		// GET: get todo from specified todo id
	e.POST("/todos", createTodo)		// POST: post new todo
	e.PATCH("/todos/:id", updateTodo)	// PATCH: update todo status
	e.DELETE("/todos/:id", deleteTodo)	// DELETE: delete specified todo

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
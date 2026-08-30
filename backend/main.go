package main

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// Todos Definition
var todos = []Todo{
	{
		Id:	1,
		Title:  "Reactを勉強する",
		Status: "done",
	},
	{
		Id:     2,
		Title:  "Goを勉強する",
		Status: "todo",
	},
}

// Next Todo Id
var nextTodoId = 3

func main() {
	// Intialize Echo
	e := echo.New()

	// CORS Middleware
	e.Use(middleware.CORS("http://localhost:5173"))

	// Handler Definition
	strictHandler := &StrictHandler{}
	serverHandler := NewStrictHandler(strictHandler, nil)

	RegisterHandlers(e, serverHandler)

	// Server Start
	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
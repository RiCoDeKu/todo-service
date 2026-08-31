package main

import (
	"context"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// ctx definition
	ctx := context.Background()

	// setting up postgreSQL DB
	db, err := NewDB(
		ctx,
		"postgres://app:password@localhost:5432/todo",
	)

	if err != nil {
		panic(err)
	}
	
	defer db.Close()

	// setting up redis client
	redisClient, err := NewRedis(
		ctx,
		"localhost:6379",
	)
	if err != nil {
		panic(err)
	}
	defer redisClient.Close()

	// Intialize Echo
	e := echo.New()

	// CORS Middleware
	e.Use(middleware.CORS("http://localhost:5173"))

	// Handler/Service/Repository/Memory Definition
	repo := NewPostgreSQLTodoRepository(db)
	service := NewTodoService(repo, redisClient)
	handler := NewTodoHandler(service)
	serverHandler := NewStrictHandler(handler, nil)
	RegisterHandlers(e, serverHandler)

	// Server Start
	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
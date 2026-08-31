package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// read info if exists .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	
	// ctx definition
	ctx := context.Background()

	// URL Settings
	databaseUrl := os.Getenv("DATABASE_URL")
	redisAddr := os.Getenv("REDIS_ADDR")

	// setting up postgreSQL DB
	db, err := NewDB(
		ctx,
		databaseUrl,
	)

	if err != nil {
		panic(err)
	}
	
	defer db.Close()

	// setting up redis client
	redisClient, err := NewRedis(
		ctx,
		redisAddr,
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
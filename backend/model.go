package main

// Type of Response / Internal Data
type Todo struct {
	ID		int		`json:"id"`
	Title	string	`json:"title"`
	Status	string	`json:"status"`
}

// for POST
type CreateTodoRequest struct {
	Title	string	`json:"title"`
}

// for PATCH
type UpdateTodoRequest struct {
	Status	string	`json:"status"`
}

type ErrorResponse struct {
	Message string	`json:"message"`
}
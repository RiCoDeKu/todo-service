package main

import (
	"context"
	"errors"
)

type StrictHandler struct{
	service *TodoService
}

func NewTodoHandler(service *TodoService) *StrictHandler {
	return &StrictHandler{
		service: service,
	}
}

var _ StrictServerInterface = (*StrictHandler)(nil)

func (h *StrictHandler) GetTodos(
	ctx context.Context,
	request GetTodosRequestObject,
) (GetTodosResponseObject, error) {
	todos, err := h.service.GetTodos(ctx)
	if err != nil {
		return GetTodos500JSONResponse{Message: "internal server error"}, nil
	}

	return GetTodos200JSONResponse(todos), nil
}

func (h *StrictHandler) GetTodo(
	ctx	context.Context,
	request GetTodoRequestObject,
) (GetTodoResponseObject, error) {
	todo, err := h.service.GetTodo(ctx, request.TodoId)
	
	if errors.Is(err, ErrTodoNotFound) {
		return GetTodo404JSONResponse{Message: "todo not found"}, nil
	}

	if err != nil {
		return GetTodo500JSONResponse{Message: "internal server error"}, nil
	}

	return GetTodo200JSONResponse(todo), nil
}

func (h *StrictHandler) CreateTodo(
	ctx context.Context,
	request CreateTodoRequestObject,
) (CreateTodoResponseObject, error) {
	if request.Body == nil {
		return CreateTodo400JSONResponse{Message: "invalid request body"}, nil
	}

	todo, err := h.service.CreateTodo(ctx, request.Body.Title)
	
	if errors.Is(err, ErrInvalidTitle){
		return CreateTodo400JSONResponse{Message: "invalid title"}, nil
	}

	if err != nil {
		return CreateTodo500JSONResponse{Message: "internal server error"}, nil
	}

	return CreateTodo201JSONResponse(todo), nil
}

func (h *StrictHandler) UpdateTodo (
	ctx context.Context,
	request UpdateTodoRequestObject,
) (UpdateTodoResponseObject, error) {

	if request.Body == nil {
		return UpdateTodo400JSONResponse{Message: "invalid request body"}, nil
	}

	todo, err := h.service.UpdateTodo(ctx, request.TodoId, TodoStatus(request.Body.Status))

	if errors.Is(err, ErrInvalidStatus) {
		return UpdateTodo400JSONResponse{Message: "invalid status"}, nil
	}

	if errors.Is(err, ErrTodoNotFound){
		return UpdateTodo404JSONResponse{Message: "todo not found"}, nil
	}

	if err != nil {
		return UpdateTodo500JSONResponse{Message: "internal server error"}, nil
	}

	return UpdateTodo200JSONResponse(todo), nil
}

func (h *StrictHandler) DeleteTodo (
	ctx context.Context,
	request DeleteTodoRequestObject,
) (DeleteTodoResponseObject, error ) {
	err := h.service.DeleteTodo(ctx, request.TodoId)

	if errors.Is(err, ErrTodoNotFound) {
		return DeleteTodo404JSONResponse{Message: "todo not found"}, nil
	}

	if err != nil {
		return DeleteTodo500JSONResponse{Message: "internal server error"}, nil
	}

	return DeleteTodo204Response{}, nil
}
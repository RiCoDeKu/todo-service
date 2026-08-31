package main

import "context"

var _ TodoRepository = (*InMemoryTodoRepository)(nil)

type TodoRepository interface {
	FindAll(ctx context.Context) ([]Todo, error)
	FindByID(ctx context.Context, id int) (Todo, bool, error)
	Create(ctx context.Context, todo Todo) (Todo, error)
	Update(ctx context.Context, todo Todo) (Todo, bool, error)
	Delete(ctx context.Context, id int) (bool, error)
}

type InMemoryTodoRepository struct {
	todos	[]Todo
	nextTodoId	int
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: []Todo{
		{Id:	1, Title:  "Reactを勉強する", 	Status: TodoStatusDone,},
		{Id:    2, Title:  "Goを勉強する",		Status: TodoStatusTodo,},
		},
		nextTodoId: 3,
	}
}

func (r *InMemoryTodoRepository) FindAll (
	ctx context.Context,
) ([]Todo, error) {
	// return todos copy cause of security.
	todos := make([]Todo, len(r.todos))
	copy(todos, r.todos)

	return todos, nil
}

func (r *InMemoryTodoRepository) FindByID (
	ctx context.Context,
	id int,
) (Todo, bool, error) {
	for _, todo := range r.todos {
		// case of found todo
		if todo.Id == id {
			return todo, true, nil
		}
	}
	// case of not found todo
	return Todo{}, false, nil
}

func (r *InMemoryTodoRepository) Create (
	ctx context.Context,
	todo Todo,
) (Todo, error) {
	todo.Id = r.nextTodoId

	r.nextTodoId++
	r.todos = append(r.todos, todo)

	return todo, nil
}

func (r *InMemoryTodoRepository) Update (
	ctx context.Context,
	todo Todo,
) (Todo, bool, error) {
	for idx, existingTodo := range r.todos {
		if existingTodo.Id == todo.Id {
			r.todos[idx].Status = todo.Status
			return r.todos[idx], true, nil
		}
	}
	//case of not found
	return Todo{}, false, nil
}

func (r *InMemoryTodoRepository) Delete (
	ctx context.Context,
	id int,
) (bool, error) {
	for idx, todo := range r.todos {
		if todo.Id == id {
			r.todos = append(r.todos[:idx], r.todos[idx+1:]...)
			return true, nil
		}
	}
	return false, nil
}
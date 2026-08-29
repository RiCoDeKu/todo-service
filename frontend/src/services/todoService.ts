import type { TodoStatus } from "../types/todo";
import type { Todo } from "../types/todo";
import type { CreateTodoRequest } from "../types/todo";
import type { UpdateTodoRequest } from "../types/todo";

export const fetchTodos = async (): Promise<Todo[]> => {
  const response = await fetch("http://localhost:8080/todos");

  if (!response.ok) {
    throw new Error("Failed to fetch todos");
  }

  return response.json();
};

export const fetchTodo = async (id: number): Promise<Todo> => {
  const response = await fetch(`http://localhost:8080/todos/${id}`);

  if (!response.ok) {
    throw new Error("Failed to fetch todo");
  }

  return response.json();
};

export const createTodo = async (input: CreateTodoRequest): Promise<Todo> => {
  const response = await fetch("http://localhost:8080/todos", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(input),
  });

  if (!response.ok) {
    throw new Error("Failed to create todo");
  }

  return response.json();
};

export const deleteTodo = async (id: number): Promise<void> => {
  const response = await fetch(`http://localhost:8080/todos/${id}`, {
    method: "DELETE",
  });

  if (!response.ok) {
    throw new Error("Failed to delete todo");
  }
};

export const updateTodo = async (
  id: number,
  input: UpdateTodoRequest,
): Promise<Todo> => {
  const response = await fetch(`http://localhost:8080/todos/${id}`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(input),
  });

  if (!response.ok) {
    throw new Error("Failed to update todo");
  }

  return response.json();
};

// Todo[] と id を受け取り、該当するTodoを返す。
export function findTodo(todos: Todo[], id: number): Todo | undefined {
  return todos.find((todo) => todo.id === id);
}

// statusで絞り込み
export function getTodosByStatus(todos: Todo[], status: TodoStatus): Todo[] {
  return todos.filter((todo) => todo.status === status);
}

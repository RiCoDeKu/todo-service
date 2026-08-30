export type TodoStatus = "todo" | "done";

export type Todo = {
  id: number;
  title: string;
  status: TodoStatus;
  description?: string;
};

export type NewTodo = Omit<Todo, "id">;

export type CreateTodoRequest = {
  title: string;
};

export type UpdateTodoRequest = {
  status: TodoStatus;
};

export type TodoStatus = "todo" | "doing" | "done";

export type Todo = {
  id: number;
  title: string;
  status: TodoStatus;
  description?: string;
};

export type NewTodo = Omit<Todo, "id">;

export type TodoFormValues = {
  title: string;
};

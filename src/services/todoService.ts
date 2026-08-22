import type { TodoStatus } from "../types/todo";
import type { Todo } from "../types/todo";
import type { NewTodo } from "../types/todo";

const todos: Todo[] = [
  {
    id: 1,
    title: "TypeScriptを勉強する",
    status: "done",
    description: "TypeScript learning is important.",
  },
  {
    id: 2,
    title: "Reactを勉強する",
    status: "todo",
    description: "React learning is important too.",
  },
  {
    id: 3,
    title: "SQLを勉強する",
    status: "todo",
  },
];

export function fetchTodos(): Promise<Todo[]> {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(todos);
      return;
    }, 1000);
  });
}

export async function fetchTodo(id: number): Promise<Todo> {
  const todos = await fetchTodos();
  const todo = findTodo(todos, id);

  if (!todo) {
    throw new Error("Todoが見つかりませんでした");
  }

  return todo;
}

export function createTodo(newTodo: NewTodo): Promise<Todo> {
  return new Promise((resolve) => {
    setTimeout(() => {
      const id =
        todos.length === 0 ? 1 : Math.max(...todos.map((todo) => todo.id)) + 1;

      const todo: Todo = {
        id,
        ...newTodo,
      };

      todos.push(todo);

      resolve(todo);
    }, 500);
  });
}

export function deleteTodo(id: number): Promise<void> {
  return new Promise((resolve) => {
    setTimeout(() => {
      const index = todos.findIndex((todo) => todo.id === id);

      if (index !== -1) {
        todos.splice(index, 1);
      }

      resolve();
    }, 500);
  });
}

export function completeTodo(id: number): Promise<Todo> {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      const todo = todos.find((todo) => todo.id === id);

      if (!todo) {
        reject(new Error("対象IDのTodoが見つかりません"));
        return;
      }

      todo.status = "done";

      resolve(todo);
    }, 500);
  });
}

// Todo[] と id を受け取り、該当するTodoを返す。
export function findTodo(todos: Todo[], id: number): Todo | undefined {
  return todos.find((todo) => todo.id === id);
}

// statusで絞り込み
export function getTodosByStatus(todos: Todo[], status: TodoStatus): Todo[] {
  return todos.filter((todo) => todo.status === status);
}

// Todo追加
export function addTodo(todos: Todo[], newTodo: NewTodo): Todo[] {
  const newId: number = todos.length + 1;

  const todo = {
    id: newId,
    ...newTodo,
  };
  return [...todos, todo];
}

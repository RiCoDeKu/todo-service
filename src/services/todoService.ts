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

export function fetchTodos(success: boolean): Promise<Todo[]> {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (success) {
        resolve(todos);
        return;
      }
      reject(new Error("Todoの取得に失敗しました"));
    }, 1000);
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

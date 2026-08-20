import { fetchTodos, getTodosByStatus, addTodo } from "./services/todoService";
import type { NewTodo, Todo } from "./types/todo";

async function main(): Promise<void> {
  try {
    console.log("取得開始");

    // fetchTodos()をawait
    const todos: Todo[] = await fetchTodos(true);
    // todo状態だけ取得
    const filteredTodos = getTodosByStatus(todos, "todo");
    console.log("todo状態のTodo:", filteredTodos);
    //新しいTodoを追加
    const newTodo: NewTodo = {
      title: "Goを勉強する",
      status: "todo",
    };
    const addedTodo = addTodo(todos, newTodo);
    console.log("追加したTodo一覧:", addedTodo);
  } catch (error) {
    if (error instanceof Error) {
      console.log("Error:", error.message);
    }
  }
  console.log("取得完了");
}

main();

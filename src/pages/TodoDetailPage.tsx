import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import type { Todo } from "../types/todo";
import { fetchTodos, findTodo } from "../services/todoService";

function TodoDetailPage() {
  const { id } = useParams();
  const [todo, setTodo] = useState<Todo | undefined>(undefined);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    async function loadTodos() {
      try {
        // Initialize each variables.
        setIsLoading(true);
        setError(null);
        setTodo(undefined);

        const todoId = Number(id);
        if (Number.isNaN(todoId)) {
          throw new Error("不正なパラメータです");
        }

        const todos = await fetchTodos(true);
        const foundTodo = findTodo(todos, todoId);

        setTodo(foundTodo);
      } catch (error) {
        if (error instanceof Error) {
          setError(error);
        }
      } finally {
        setIsLoading(false);
      }
    }
    loadTodos();
  }, [id]);

  if (isLoading) return <p>読み込み中</p>;
  if (error) return <p>エラー: {error.message}</p>;
  if (!todo) return <p>Todoが見つかりません</p>;
  return (
    <>
      <p>タイトル：{todo.title}</p>
      <p>ステータス：{todo.status === "done" ? "完了" : todo.status}</p>
      {todo.description && <p>説明：{todo.description}</p>}
    </>
  );
}

export default TodoDetailPage;

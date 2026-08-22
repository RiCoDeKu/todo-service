import { useParams } from "react-router-dom";
import { fetchTodo } from "../services/todoService";
import { useQuery } from "@tanstack/react-query";

function TodoDetailPage() {
  const { id } = useParams();

  const todoId = Number(id);
  const isValidId = !Number.isNaN(todoId);

  const {
    data: todo,
    isPending,
    error,
  } = useQuery({
    queryKey: ["todos", todoId],
    queryFn: () => fetchTodo(todoId),
    enabled: isValidId,
  });

  if (!isValidId) return <p>不正なIDです</p>;
  if (isPending) return <p>読み込み中</p>;
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

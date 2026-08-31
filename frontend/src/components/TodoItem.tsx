import { Link } from "react-router-dom";
import type { Todo } from "../types/todo";

type TodoItemProps = {
  todo: Todo;
  onDelete: (id: number) => void;
  onComplete: (id: number) => void;
  isDisable: boolean;
};

function TodoItem({ todo, onDelete, onComplete, isDisable }: TodoItemProps) {
  return (
    <>
      <p>タイトル：{todo.title}</p>
      <p>ステータス：{todo.status}</p>
      <button>
        <Link to={`/todos/${todo.id}`}>詳細</Link>
      </button>
      <button
        onClick={() => {
          onComplete(todo.id);
        }}
		disabled={isDisable}
      >
        完了にする
      </button>
      <button
        onClick={() => {
          onDelete(todo.id);
        }}
		disabled={isDisable}
      >
        削除
      </button>
    </>
  );
}

export default TodoItem;

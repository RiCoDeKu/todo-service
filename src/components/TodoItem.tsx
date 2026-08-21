import { Link } from "react-router-dom";
import type { Todo } from "../types/todo";

type TodoItemProps = {
  todo: Todo;
  onDelete: (id: number) => void;
  onComplete: (id: number) => void;
};

function TodoItem({ todo, onDelete, onComplete }: TodoItemProps) {
  return (
    <>
      <p>タイトル：{todo.title}</p>
      <button>
        <Link to={`/todos/${todo.id}`}>詳細</Link>
      </button>
      <button
        onClick={() => {
          onComplete(todo.id);
        }}
      >
        完了にする
      </button>
      <button
        onClick={() => {
          onDelete(todo.id);
        }}
      >
        削除
      </button>
    </>
  );
}

export default TodoItem;

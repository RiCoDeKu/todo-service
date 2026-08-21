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
      <p>ステータス：{todo.status === "done" ? "完了" : todo.status}</p>
      {todo.description && <p>説明：{todo.description}</p>}
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

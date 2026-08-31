import { useAtom } from "jotai";
import { todoFilterAtom } from "../atoms/todoFilter";
import { ResetTodoFilterButton } from "./ResetTodoFilterButton";

export function TodoFilter() {
  const [statusFilter, setStatusFilter] = useAtom(todoFilterAtom);

  return (
    <div>
      <button
        onClick={() => {
          setStatusFilter("all");
        }}
        disabled={statusFilter === "all"}
      >
        すべて
      </button>
      <button
        onClick={() => {
          setStatusFilter("todo");
        }}
        disabled={statusFilter === "todo"}
      >
        未着手
      </button>
      <button
        onClick={() => {
          setStatusFilter("doing");
        }}
        disabled={statusFilter === "doing"}
      >
        進行中
      </button>
      <button
        onClick={() => {
          setStatusFilter("done");
        }}
        disabled={statusFilter === "done"}
      >
        完了
      </button><br/>
      <ResetTodoFilterButton />
    </div>
  );
}

export default TodoFilter;

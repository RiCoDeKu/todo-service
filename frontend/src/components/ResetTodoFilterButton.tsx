import { useSetAtom } from "jotai";
import { todoFilterAtom } from "../atoms/todoFilter";

export function ResetTodoFilterButton() {
  const setStatusFilter = useSetAtom(todoFilterAtom);

  return <button onClick={() => setStatusFilter("all")}>フィルターをリセット</button>;
}

export default ResetTodoFilterButton;

import type { TodoFormValues } from "../types/todo";
import { Link } from "react-router-dom";
import TodoItem from "../components/TodoItem";
import {
  fetchTodos,
  createTodo,
  deleteTodo,
  completeTodo,
} from "../services/todoService";
import TodoFilter from "../components/TodoFilter";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { useAtomValue } from "jotai";
import { todoFilterAtom } from "../atoms/todoFilter";
import { useForm } from "react-hook-form";

function TodoListPage() {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<TodoFormValues>();
  const statusFilter = useAtomValue(todoFilterAtom);

  const queryClient = useQueryClient();

  const createTodoMutation = useMutation({
    mutationFn: createTodo,

    onSuccess: () => {
      reset();

      queryClient.invalidateQueries({
        queryKey: ["todos"],
      });
    },
  });

  const deleteTodoMutation = useMutation({
    mutationFn: deleteTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["todos"],
      });
    },
  });

  const completeTodoMutation = useMutation({
    mutationFn: completeTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["todos"],
      });
    },
  });

  function handleAddTodo(data: TodoFormValues) {
    createTodoMutation.mutate({
      title: data.title,
      status: "todo",
    });
  }

  function handleDeleteTodo(id: number) {
    deleteTodoMutation.mutate(id);
  }

  function handleCompleteTodo(id: number) {
    completeTodoMutation.mutate(id);
  }

  const {
    data: todos,
    isPending,
    error,
  } = useQuery({
    queryKey: ["todos"],
    queryFn: fetchTodos,
    staleTime: 60_000,
  });

  function filterTodos() {
    if (!todos) {
      return null;
    }

    const filteredTodos =
      statusFilter === "all"
        ? todos
        : todos.filter((todo) => todo.status === statusFilter);

    if (filteredTodos.length === 0) {
      return <p>該当するTodoはありません</p>;
    }
    return filteredTodos.map((todo) => (
      <TodoItem
        key={todo.id}
        todo={todo}
        onDelete={handleDeleteTodo}
        onComplete={handleCompleteTodo}
        isDisable={
          deleteTodoMutation.isPending || completeTodoMutation.isPending
        }
      />
    ));
  }

  function renderTodos() {
    if (isPending) {
      return <p>読み込み中</p>;
    }

    if (error) {
      return <p>エラーが発生しました：{error.message}</p>;
    }

    return (
      <>
        <TodoFilter />
        {filterTodos()}
        <br />
        <form onSubmit={handleSubmit(handleAddTodo)}>
          <input
            {...register("title", {
              required: "タイトルは必須です",
              maxLength: {
                value: 100,
                message: "タイトルは100文字以内で入力してください",
              },
              validate: (value) =>
                value.trim() !== "" || "タイトルを正しく入力してください",
            })}
          />
          <button type="submit" disabled={createTodoMutation.isPending}>
            {createTodoMutation.isPending ? "追加中..." : "追加"}
          </button>
        </form>
        {errors.title && <p>{errors.title.message}</p>}
        {createTodoMutation.isError && <p>追加に失敗しました</p>}
        <br />
      </>
    );
  }

  return (
    <div>
      <h1>Todo App</h1>
      {renderTodos()}
      <Link to={"/"}>Homeへ戻る</Link>
    </div>
  );
}

export default TodoListPage;

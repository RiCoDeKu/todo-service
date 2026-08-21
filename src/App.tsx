import { useState, useEffect } from "react";
import TodoItem from "./components/TodoItem";
import type { Todo } from "./types/todo";
import { fetchTodos } from "./services/todoService";

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [title, setTitle] = useState("");
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  function handleAddTodo() {
    if (title.trim() === "") {
      return;
    }

    setTodos((prevTodos) => {
      const newId: number =
        prevTodos.length === 0
          ? 1
          : Math.max(...prevTodos.map((todo) => todo.id)) + 1;

      const newTodo: Todo = {
        id: newId,
        title,
        status: "todo",
      };
      return [...prevTodos, newTodo];
    });

    setTitle("");
  }

  function handleDeleteTodo(id: number) {
    setTodos((prevTodos) => prevTodos.filter((todo) => todo.id !== id));
  }

  function handleCompleteTodo(id: number) {
    setTodos((prevTodos) =>
      prevTodos.map((todo) => {
        if (todo.id === id) {
          return {
            ...todo,
            status: "done",
          };
        }
        return todo;
      }),
    );
  }

  useEffect(() => {
    async function loadTodos() {
      try {
        const result = await fetchTodos(true);
        setTodos(result);
      } catch (error) {
        if (error instanceof Error) {
          setError(error);
        }
      } finally {
        setIsLoading(false);
      }
    }
    loadTodos();
  }, []);

  function renderTodos() {
    if (isLoading) {
      return <p>読み込み中</p>;
    }

    if (error) {
      return <p>エラーが発生しました：{error.message}</p>;
    }

    return (
      <>
        {todos.length === 0 ? (
          <p>todoがありません</p>
        ) : (
          <>
            {todos.map((todo) => (
              <TodoItem
                key={todo.id}
                todo={todo}
                onDelete={handleDeleteTodo}
                onComplete={handleCompleteTodo}
              />
            ))}
          </>
        )}
        <br />
        <input
          value={title}
          onChange={(event) => {
            setTitle(event.target.value);
          }}
        />
        <button onClick={handleAddTodo}>追加</button>
      </>
    );
  }

  return (
    <div>
      <h1>Todo App</h1>
      {renderTodos()}
    </div>
  );
}

export default App;

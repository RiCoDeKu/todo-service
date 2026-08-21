import { Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage";
import TodoListPage from "./pages/TodoListPage";
import TodoDetailPage from "./pages/TodoDetailPage";

function App() {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/todos" element={<TodoListPage />} />
      <Route path="/todos/:id" element={<TodoDetailPage />} />
    </Routes>
  );
}

export default App;

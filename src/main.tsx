import { StrictMode } from "react"; // 開発時に問題を見つけやすくするReactの機能
import { createRoot } from "react-dom/client"; // ReactをブラウザのDOMへ接続する関数
import { BrowserRouter } from "react-router-dom"; // 現在のブラウザURLを監視し、React Routerで扱えるようにする
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import App from "./App"; // 自分で作ったReactコンポーネント

const queryClient = new QueryClient();

createRoot(document.getElementById("app")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </QueryClientProvider>
  </StrictMode>,
);

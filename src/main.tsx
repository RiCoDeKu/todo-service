import { StrictMode } from "react"; // 開発時に問題を見つけやすくするReactの機能
import { createRoot } from "react-dom/client"; // ReactをブラウザのDOMへ接続する関数
import App from "./App"; // 自分で作ったReactコンポーネント

createRoot(document.getElementById("app")!).render(
  <StrictMode>
    <App />
  </StrictMode>,
);

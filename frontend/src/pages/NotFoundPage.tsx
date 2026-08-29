import { Link } from "react-router-dom";

function NotFoundPage() {
  return (
    <>
      <h1>404 Not Found</h1>
      <p>ページが見つかりません</p>
      <Link to="/">Homeへ戻る</Link>
    </>
  );
}

export default NotFoundPage;

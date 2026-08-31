import { Link } from "react-router-dom";

function HomePage() {
  return (
    <>
      <h1>Home</h1>
      <Link to={"/todos"}>Todo一覧へ</Link>
    </>
  );
}

export default HomePage;

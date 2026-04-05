import { useSearchParams } from "react-router";

const LoginPage = () => {
  const [searchParams] = useSearchParams();
  const token = searchParams.get("token");

  if (!token) {
    return <div>No token provided</div>;
  }

  return <div>Login token: {token}</div>;
};

export default LoginPage;

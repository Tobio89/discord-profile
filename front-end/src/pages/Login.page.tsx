import { useQuery } from "@tanstack/react-query";
import { useSearchParams } from "react-router";

async function validateToken(token: string) {
  const request = new Request(`http://localhost:4455/validate-token/${token}`, {
    method: "POST",
  });

  try {
    const response = await fetch(request);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }
    const result = await response.json();
    return result;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

const useLogin = (token: string) => {
  const query = useQuery({
    queryKey: ["login", token],
    queryFn: async () => validateToken(token),
  });
  return query;
};

const LoginPart = ({ token }: { token: string }) => {
  const { data, isLoading, isError, error } = useLogin(token);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return (
      <div>
        Error: {error instanceof Error ? error.message : "Unknown error"}
      </div>
    );
  }

  return <div>Login successful! User ID: {JSON.stringify(data)}</div>;
};

const LoginPage = () => {
  const [searchParams] = useSearchParams();
  const token = searchParams.get("token");

  if (!token) {
    return <div>No token provided</div>;
  }

  return <LoginPart token={token} />;
};

export default LoginPage;

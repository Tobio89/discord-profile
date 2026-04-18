import { useMutation, useQuery } from "@tanstack/react-query";
import React from "react";
import { useSearchParams } from "react-router";

async function validateToken(token: string) {
  const request = new Request(`http://localhost:4455/validate-token/${token}`, {
    method: "POST",
    credentials: "include",
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

async function checkJWT() {
  const request = new Request(`http://localhost:4455/check-token`, {
    method: "GET",
    credentials: "include",
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

const useJWTCheck = () => {
  const query = useMutation({
    mutationKey: ["jwt-check"],
    mutationFn: async (responseHandler: (result: unknown) => void) => {
      responseHandler("Checking JWT...");
      const result = await checkJWT();
      responseHandler(result);
    },
  });
  return query;
};

const LoginPart = ({ token }: { token: string }) => {
  const { data, isLoading, isError, error } = useLogin(token);
  const { mutate: checkJWT } = useJWTCheck();

  const [response, setResponse] = React.useState<string | null>(null);

  const handleJWTCheck = () => {
    checkJWT((result) => {
      setResponse(JSON.stringify(result));
    });
  };

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

  return (
    <div>
      <div>Login successful! User ID: {JSON.stringify(data)}</div>
      <button onClick={() => handleJWTCheck()}>Check JWT</button>
      <div>JWT check response: {response}</div>
    </div>
  );
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

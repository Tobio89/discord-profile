import { routeLoader$ } from "@builder.io/qwik-city";

export const useValidateToken = routeLoader$(
  async ({ url, redirect, cookie }) => {
    const token = url.searchParams.get("token");

    if (!token) {
      throw redirect(302, "/login");
    }

    const request = new Request(
      `http://localhost:4455/validate-token/${token}`,
      {
        method: "POST",
        credentials: "include",
      },
    );

    const res = await fetch(request);

    if (!res.ok) {
      throw redirect(302, "/login?error=invalid-token");
    }

    const data = await res.json();
    cookie.set("session", data.data.jwt, {
      path: "/",
      httpOnly: true,
      secure: true,
      sameSite: "lax",
    });

    throw redirect(302, "/dashboard");
  },
);

export default () => {
  return <p>Processing login...</p>;
};

import { component$, Slot } from "@builder.io/qwik";
import { routeLoader$ } from "@builder.io/qwik-city";

export const useAuthGuard = routeLoader$(async ({ cookie, redirect }) => {
  const session = cookie.get("session")?.value;

  if (!session) {
    throw redirect(302, "/login");
  }

  const res = await fetch("http://localhost:4455/check-token", {
    headers: {
      Cookie: `session=${session}`,
    },
  });

  if (!res.ok) {
    throw redirect(302, "/login");
  }

  return res.json();
});

export default component$(() => {
  useAuthGuard(); // runs automatically for all child routes
  return <Slot />;
});

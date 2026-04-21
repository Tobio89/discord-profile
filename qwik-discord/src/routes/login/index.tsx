import { component$ } from "@builder.io/qwik";
import { useLocation, type DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  const loc = useLocation();
  const errorMessage = loc.url.searchParams.get("error");
  return (
    <>
      <nav class="nav-bar">
        <a href="/">Home</a>
      </nav>
      <main class="container">
        <h1>Discord Profile: Login</h1>
        <div>
          {errorMessage === "missing-token" && (
            <p>
              Error: No token! Please use the <code>/login</code> command in the
              Discord server to get a valid login link.
            </p>
          )}
          {errorMessage === "invalid-token" && (
            <p>
              Error: Invalid token! Please use the <code>/login</code> command
              in the Discord server to get a valid login link.
            </p>
          )}
          {!errorMessage && (
            <p>
              Login with Discord to view your profile information and manage
              your account settings.
            </p>
          )}
        </div>
      </main>
    </>
  );
});

export const head: DocumentHead = {
  title: "Discord Profile: Login",
};

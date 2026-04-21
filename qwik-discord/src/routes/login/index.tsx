import { component$ } from "@builder.io/qwik";
import { useLocation, type DocumentHead } from "@builder.io/qwik-city";
import ErrorMessage from "~/components/ErrorMessage";

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
          <ErrorMessage errorCode={errorMessage} />
          <p>
            Login with Discord to view your profile information and manage your
            account settings.
          </p>
        </div>
      </main>
    </>
  );
});

export const head: DocumentHead = {
  title: "Discord Profile: Login",
};

import { component$ } from "@builder.io/qwik";
import { type DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <>
      <nav class="nav-bar">
        <a href="/">Home</a>
      </nav>
      <main class="container">
        <h1>Discord Profile: User Page</h1>
        <div>
          <p>If you are here, you've probably logged in!</p>
        </div>
      </main>
    </>
  );
});
export const head: DocumentHead = {
  title: "Discord Profile: DASH",
};

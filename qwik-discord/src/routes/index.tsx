import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <>
      <nav class="nav-bar">
        <a href="/">Home</a>
      </nav>
      <main>
        <h1>Discord Profile</h1>
        <div>
          <p>
            It connects to a Discord server via a Bot to provide a detailed
            profile page for users!
          </p>
          <p>
            To log in, visit the discord server, and use <code>/login</code> to
            get a login link.
          </p>
        </div>
      </main>
    </>
  );
});

export const head: DocumentHead = {
  title: "Discord Profile",
  meta: [
    {
      name: "description",
      content: "A profile page connected to Discord",
    },
  ],
};

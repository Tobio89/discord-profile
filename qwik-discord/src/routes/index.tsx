import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <>
      <h1>Discord Profile</h1>
      <div>
        <p>
          To log in, visit the discord server, and use <code>/login</code> to
          get a login link.
        </p>
      </div>
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

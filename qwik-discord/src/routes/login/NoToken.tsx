import { component$ } from "@builder.io/qwik";
import { type DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <>
      <div>
        <p>Looks like you've got no token.</p>
        <p>
          Go to the discord server and use <code>/login</code> to start logging
          in.
        </p>
      </div>
    </>
  );
});

export const head: DocumentHead = {
  title: "Discord Profile: Login",
};

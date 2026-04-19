import { component$ } from "@builder.io/qwik";
import { type DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <>
      <h1>Discord Profile: Login</h1>
      <div>
        Login with Discord to view your profile information and manage your
        account settings.
      </div>
    </>
  );
});

export const head: DocumentHead = {
  title: "Discord Profile: Login",
};

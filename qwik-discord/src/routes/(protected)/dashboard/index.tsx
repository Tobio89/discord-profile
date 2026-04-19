import { component$ } from "@builder.io/qwik";
import { type DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <div>
      <h1>Hello!</h1>
      <p>If you are here, you logged in</p>
      {/* <p>{JSON.stringify(session)}</p> */}
    </div>
  );
});
export const head: DocumentHead = {
  title: "Discord Profile: DASH",
};

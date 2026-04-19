import { component$ } from "@builder.io/qwik";
import { useLocation, type DocumentHead } from "@builder.io/qwik-city";
import NoToken from "./NoToken";
import LoginWithToken from "./LoginWithToken";

export default component$(() => {
  const loc = useLocation();
  const loginToken = loc.url.searchParams.get("token");

  if (!loginToken) {
    return (
      <>
        <h1>Discord Profile: Login</h1>
        <NoToken />
      </>
    );
  }

  return (
    <>
      <h1>Discord Profile: Login</h1>
      <div>
        <p>You're trying to login - your token is: {loginToken}</p>
      </div>
      <LoginWithToken token={loginToken} />
    </>
  );
});

export const head: DocumentHead = {
  title: "Discord Profile: Login",
};

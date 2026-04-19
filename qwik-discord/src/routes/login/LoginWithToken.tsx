import { component$, Resource, useResource$, $ } from "@builder.io/qwik";
import { type DocumentHead } from "@builder.io/qwik-city";

interface Props {
  token: string;
}

export default component$<Props>((props) => {
  const validateToken$ = $(async function validateToken(token: string) {
    const request = new Request(
      `http://localhost:4455/validate-token/${token}`,
      {
        method: "POST",
        credentials: "include",
      },
    );

    try {
      const response = await fetch(request);
      if (!response.ok) {
        throw new Error(`Response status: ${response.status}`);
      }
      const result = await response.json();
      return result;
    } catch (error) {
      console.error(error);
      throw error;
    }
  });

  const tokenResource = useResource$<string[]>(({ track, cleanup }) => {
    // We need a way to re-run fetching data whenever the `github.org` changes.
    // Use `track` to trigger re-running of this data fetching function.
    track(() => props.token);

    // A good practice is to use `AbortController` to abort the fetching of data if
    // new request comes in. We create a new `AbortController` and register a `cleanup`
    // function which is called when this function re-runs.
    const controller = new AbortController();
    cleanup(() => controller.abort());

    // Fetch the data and return the promises.
    return validateToken$(props.token);
  });

  return (
    <>
      <div>
        <Resource
          value={tokenResource}
          onPending={() => <p>Validating token...</p>}
          onResolved={(data) => (
            <div>
              <p>Token is valid! Your data: {JSON.stringify(data)}</p>
            </div>
          )}
          onRejected={(error) => (
            <div>
              <p>Token validation failed: {error.message}</p>
            </div>
          )}
        />
      </div>
    </>
  );
});

export const head: DocumentHead = {
  title: "Discord Profile: Login",
};

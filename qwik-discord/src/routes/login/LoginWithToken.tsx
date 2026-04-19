import { component$, Resource, useResource$, $ } from "@builder.io/qwik";
import { type DocumentHead } from "@builder.io/qwik-city";

interface Props {
  token: string;
}

export default component$<Props>((props) => {
  const validateToken$ = $(async function validateToken(
    token: string,
    controller: AbortController,
  ) {
    const request = new Request(
      `http://localhost:4455/validate-token/${token}`,
      {
        method: "POST",
        credentials: "include",
        signal: controller.signal,
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
    track(() => props.token);
    const controller = new AbortController();
    cleanup(() => controller.abort());

    // Fetch the data and return the promises.
    return validateToken$(props.token, controller);
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

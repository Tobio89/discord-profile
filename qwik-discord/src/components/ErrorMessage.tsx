import { component$ } from "@builder.io/qwik";

interface Props {
  errorCode: string | null;
}

export default component$<Props>((props) => {
  if (props.errorCode === "missing-token") {
    return <p class="error-message">Error: No token!</p>;
  }

  if (props.errorCode === "invalid-token") {
    return <p class="error-message">Error: Invalid token!</p>;
  }

  return null;
});

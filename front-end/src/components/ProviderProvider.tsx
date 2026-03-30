import type { PropsWithChildren } from "react";
import { BrowserRouter } from "react-router";

const ProviderProvider = ({ children }: PropsWithChildren) => {
  return <BrowserRouter>{children}</BrowserRouter>;
};

export default ProviderProvider;

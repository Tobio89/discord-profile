import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import App from "./App.tsx";
import ProviderProvider from "./components/ProviderProvider.tsx";

import "./index.css";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ProviderProvider>
      <App />
    </ProviderProvider>
  </StrictMode>,
);

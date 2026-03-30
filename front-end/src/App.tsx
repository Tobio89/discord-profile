import { useRoutes } from "react-router";
import UserPage from "./pages/User.page";

const router = [
  {
    path: "/:userID",
    element: <UserPage />,
  },
  {
    path: "*",
    element: <UserPage />,
  },
];

function App() {
  const element = useRoutes(router);

  return <>{element}</>;
}

export default App;

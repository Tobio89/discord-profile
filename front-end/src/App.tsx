import { useRoutes } from "react-router";
import UserPage from "./pages/User.page";
import LoginPage from "./pages/Login.page";

const router = [
  {
    path: "/login",
    element: <LoginPage />,
  },
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

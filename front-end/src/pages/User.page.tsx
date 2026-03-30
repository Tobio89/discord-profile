import { useParams } from "react-router";

const UserPage = () => {
  const { userID } = useParams();

  if (!userID) {
    return <div>No user ID provided</div>;
  }

  return <div>User ID: {userID}</div>;
};

export default UserPage;

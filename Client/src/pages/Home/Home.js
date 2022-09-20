import { useAuth } from "../../hooks/useAuth";
import { Navigate } from "react-router-dom";

export default function Home () {

  const { user } = useAuth();

  console.log(user)

  if (!user) {
    return <Navigate to="/login" />;
  }

  return (
    <h1>Home</h1>
  )
} 
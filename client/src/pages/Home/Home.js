import { Navigate } from "react-router-dom";
import Cookies from "js-cookie";

export default function Home () {

  var user = Cookies.get("token")

  console.log(user)

  if (!user) {
    return <Navigate to="/login" />;
  }

  return (
    <div>
      <header>
        
      </header>
      <h1>Home</h1>
    </div>

  )
} 
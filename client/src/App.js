import "./App.css";
import { Routes,Route } from "react-router-dom";
import Login from "./pages/Login/Login";
import Hero from "./pages/Hero/Hero";
import Home from "./pages/Home/Home";
import Logout from "./pages/Login/Logout";
import SignUp from "./pages/SignUp/SignUp"
import AddFail from "./pages/AddFail/AddFail";
import GetFailList from "./pages/GetFailList.js/GetFailList";

function App() {
  return (
        <Routes>
          <Route path="/" element={<Hero />}/>
          <Route path="signup" element={<SignUp />}/>
          <Route path="logout" element={<Logout/>}/>
          <Route path="login" element={<Login/>} />
          <Route path="home" element={<Home/>}>
            <Route path="add" element={<AddFail/>}/>
            <Route path="list" element={<GetFailList/>}/>
          </Route>
        </Routes>
  );
}

export default App;

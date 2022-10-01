import "./App.css";
import { Routes,Route } from "react-router-dom";
import Login from "./pages/Login/Login";
import Hero from "./pages/Hero/Hero";
import Home from "./pages/Home/Home";
import Logout from "./pages/Login/Logout";

function App() {
  return (
        <Routes>
          <Route path="/" element={<Hero />}/>
          <Route path="/logout" element={<Logout/>}/>
          <Route path="/login" element={<Login/>} />
          <Route path="/home" element={<Home/>} />
        </Routes>
  );
}

export default App;

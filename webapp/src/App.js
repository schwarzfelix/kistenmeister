import './App.css';
import { Route, Routes } from 'react-router-dom';
import Aktivierung from './pages/Aktivierung';
import Comments from './pages/Comments';
import Details from './pages/Details';
import Home from './pages/Home';
import Images from './pages/Images';
import List from './pages/List';
import NewBox from './pages/NewBox';
import QRCode from './pages/QRCode';
import Registrierung from './pages/Registrierung';
import React, { useEffect, useState } from 'react';
import Profile from './pages/Profile';

export const UserContext = React.createContext();

function App() {

  const [user, setUser] = useState({
    id: localStorage.getItem("id"),
    name: localStorage.getItem("name"),
    email: localStorage.getItem("email"),
    token: localStorage.getItem("token")
  });

  const userContextValue = { user, setUser };


  useEffect(() => {
    console.log("App.js useEffect");
    console.log(user);
  }, [user]);

  return (
    <UserContext.Provider value={userContextValue}>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/box/:id" element={<Details />} />
        <Route path="/box/:id/qr" element={<QRCode />} />
        <Route path="/box/:id/comments" element={<Comments />} />
        <Route path="/box/:id/images" element={<Images />} />
        <Route path="/new" element={<NewBox />} />
        <Route path="/profile" element={<Profile />} />
        <Route path="/register" element={<Registrierung />} />
        <Route path="/list" element={<List />} />
        <Route path="/aktivierung" element={<Aktivierung />} />
        <Route path="*" element={<h1>Not Found</h1>} />
      </Routes>
    </UserContext.Provider>
  );
}

export default App;

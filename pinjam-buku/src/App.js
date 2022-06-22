import React from 'react'
import { Routes, Route, Link } from "react-router-dom";
import Register from "./account/Register";
import Home from './user/Home';
import Login from './account/Login'

function App() {
  return (
    <div className="App">
      <Routes>        
        <Route index element={<Home />} />
        <Route path="/daftar" element={<Register />} />
        <Route path='/masuk' element={<Login />} />
      </Routes>
    </div>
  );
}

export default App;
import React from 'react'
import { Routes , Route } from 'react-router-dom'
import Login from './components/Login'
import SignUp from './components/Signup'
import Dashboard from './components/Dashboard'

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/Signup" element={<SignUp />} />
        <Route path="/Dashboard" element={<Dashboard />} />
      </Routes>
    </div>
  )
}
export default App
import React, { useState, Fragment, useEffect } from "react";
import { Link, useNavigate} from "react-router-dom";
import axios from 'axios'
import Header from '../components/Header'

function Login (props) {
  const [user, setUser] = useState({email: "", password:""})
  const apiUrl = "/https://api-dev.pinjambuku.me/api/v1/auth/login"
  const Logins = (e) => {
    e.preventDefault();    
            debugger;   
            const data = { email: user.email, password: user.password };    
            axios.post(apiUrl, data)    
            .then((result) => {    
                debugger;  
                console.log(result.data);   
                const serializedState = JSON.stringify(result.data.UserDetails);  
               var a= localStorage.setItem('myData', serializedState);   
                console.log("A:",a)  
                const user =result.data.UserDetails;  
                console.log(result.data.message);  
                if (result.data.status == '200')    
                    props.history.push('/Dashboard')    
                else    
                alert('Invalid User');    
            })         
  }

  const onChange = (e) => {
    e.persist();    
    debugger;    
    setUser({...user, [e.target.name]: e.target.value});
  }
  // const[email, setEmail] = useState("");
  // const[password, setPassword] = useState("");
  // const[redirect, setRedirect] = useState(false);
  // const[error, setError] = useState("");

  // const onChangeEmail = (e) => {
  //   const value = e.target.value;
  //   setEmail(value);
  // }

  // const onChangePassword = (e) => {
  //   const value = e.target.value;
  //   setPassword(value);
  // }

  // const SubmitLogin = async (e) => {
  //   const navigate = useNavigate()
  //   const data = {
  //     email: email,
  //     password: password
  //   }
   
  //     const response = await axios.post('https://api-dev.pinjambuku.me/api/auth/login', data)
  //     console.log(response);
  //     if (response){
  //       navigate("/daftar")
  //     }}
  

  return (
    <Fragment>
        <Header />
        <div style={{ marginTop: "50px" }}>
            <div className="container">
                <div className="row justify-content-end">
                    <div className="col-md-5">
                        <div className="card p-4 shadow rounded">
                            <div className="card-body">
                              <div className="row">
                                  <div className="col-md-6">
                                  <h1 className="mt-3 mb-3">Masuk</h1>
                                  </div>
                                  <div className="col-md-6 text-end mt-4">
                                  <Link to="/daftar" className="text-decoration-none">Daftar</Link>
                                  </div>
                              </div>
                              
                              <form onSubmit={Logins} class="user">
                                <div className="form-group">
                                    <label className="mb-2">Email</label>
                                    <input type="email" className="form-control" value={user.email} onChange={ onChange }  
                                    name="email" id="email" aria-describedby="emailHelp" placeholder="Enter Email"/>
                                </div>
                                <div className="form-group">
                                    <label className="mb-2 mt-2">Password</label>
                                    <input type="password" className="form-control" value={user.password} onChange={ onChange }  
                                    name="password" id="DepPasswordartment" placeholder="Password"/>
                                </div>
                                 
                                <p className="forgot-password text-right mt-4 text-end">
                                    <a href="#" className="text-decoration-none" style={{color: "red"}}>
                                        Lupa Kata Sandi?
                                    </a></p>
                                
                                <div className="d-grid">
                                    <button type="submit" className="btn btn-primary btn-block mt-3 mb-4">Masuk</button>
                                </div>
                              </form>
                          </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </Fragment>
  );
}

export default Login;
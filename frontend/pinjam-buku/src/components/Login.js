import React, { useState, Fragment } from "react";
import { Link, NavLink} from "react-router-dom";
import axios from 'axios'

const Login = () => {
  const[email, setEmail] = useState("");
  const[password, setPassword] = useState("");
  const[redirect, setRedirect] = useState(false);
  const[error, setError] = useState("");

  const onChangeEmail = (e) => {
    const value = e.target.value;
    setEmail(value);
  }

  const onChangePassword = (e) => {
    const value = e.target.value;
    setPassword(value);
  }

  const submitLogin = (e) => {
    const data = {
      email: email,
      password: password
    }
    axios.post('https://api-dev.pinjambuku.me/', data)
    .then(result => {
      if(result){
        localStorage.setItem('token', result.data.token)
        setRedirect(true)
      }
    })
  
  }

  return (
    <Fragment>
      { 
        redirect ? ( 
        <NavLink to="/Dashboard" />)
      :(
        null)
      }
      <div style={{ marginTop: "100px" }}>
        <div className="container">
          <div className="row justify-content-end">
            <div className="col-md-5">
              <div className="card p-4">
                {

                  error ? (
                    <div className="alert alert-danger">
                      {error}
                    </div>
                  ) : null
                }
                <div className="card-body">
                  <div className="row">
                    <div className="col-md-6">
                      <h1 className="mt-3 mb-3">Masuk</h1>
                    </div>
                    <div className="col-md-6 text-end mt-4">
                      <Link to="/Signup">Daftar?</Link>
                    </div>
                  </div>

                  <div className="form-group">
                    <label className="mb-2">Email</label>
                    <input type="text" className="form-control" onChange={onChangeEmail} />
                  </div>
                  <div className="form-group">
                    <label className="mb-2 mt-2">Password</label>
                    <input type="password" className="form-control" onChange={onChangePassword} />
                  </div>
                  <p className="forgot-password text-right mt-2"><a href="#">Lupa Kata Sandi?</a></p>
                  <div className="d-grid">
                    <button className="btn btn-primary btn-block mt-3 mb-4" onClick={submitLogin}>Masuk</button>
                  </div>
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


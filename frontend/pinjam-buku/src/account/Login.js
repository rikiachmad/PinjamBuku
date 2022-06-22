import React, { useState, Fragment } from "react";
import { Link, useNavigate} from "react-router-dom";
import axios from 'axios'
import Header from '../components/Header'

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

  const SubmitLogin = async (e) => {
    const navigate = useNavigate()
    const data = {
      email: email,
      password: password
    }
   
      const response = await axios.post('https://api-dev.pinjambuku.me/api/auth/login', data)
      console.log(response);
      if (response){
        navigate("/daftar")
      }
  
}
  return (
    <Fragment>
        <Header />
        <div style={{ marginTop: "50px" }}>
            <div className="container">
                <div className="row justify-content-end">
                    <div className="col-md-5">
                        <div className="card p-4 shadow rounded">
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
                                <Link to="/daftar" className="text-decoration-none">Daftar</Link>
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
                            <p className="forgot-password text-right mt-4 text-end">
                                <a href="#" className="text-decoration-none" style={{color: "red"}}>
                                    Lupa Kata Sandi?
                                </a></p>
                            <div className="d-grid">
                                <button className="btn btn-primary btn-block mt-3 mb-4" onClick={SubmitLogin}>Masuk</button>
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
import React, { useState } from "react";
import { Link } from "react-router-dom";

const Signup = () => {
  const[email, setEmail] = useState("");
  const[password, setPassword] = useState("");

  const onChangeEmail = (e) => {
    const value = e.target.value;
    setEmail(value);
  }

  const onChangePassword = (e) => {
    const value = e.target.value;
    setPassword(value);
  }

  return (
    <div style={{ marginTop: "100px" }}>
      <div className="container">
        <div className="row justify-content-end">
          <div className="col-md-5">
            <div className="card p-4">
              <div className="card-body">
              <div className="row">
                  <div className="col-md-6">
                    <h1 className="mt-3 mb-3">Daftar</h1>
                  </div>
                  <div className="col-md-6 text-end mt-4">
                    <Link to="/">Masuk?</Link>
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
                <div className="d-grid">
                  <button className="btn btn-primary btn-block mt-3 mb-4">Daftar</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Signup;
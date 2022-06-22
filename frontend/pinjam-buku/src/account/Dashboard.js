import React, { useState, useEffect } from 'react'  
  
function Dashboard() {  
    const [user, setuser] = useState({ user: '', password: '' });  
    useEffect(() => {  
        var a = localStorage.getItem('myData');  
        var b = JSON.parse(a);  
        console.log(b.userName);  
        setuser(b)  
        console.log(user.userName)  
  
    }, []);  
    return (  
        <>  
            <div class="col-sm-12 btn btn-primary">  
                Dashboard  
        </div>  
            <h1>Welcome :{user.userName}</h1>  
        </>  
    )  
}  
  
export default Dashboard  
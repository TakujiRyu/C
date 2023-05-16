function signUp() {
    var _data = {
        firstname : document.getElementById("fname").value,
        lastname : document.getElementById("lname").value,
        email : document.getElementById("email").value,
        password : document.getElementById("pw1").value,
        pw : document.getElementById("pw2").value
    }
    if (_data.password !== _data.pw) {
        alert("PASSWORD doesn't match!")
        return
    }
    fetch('/signup', {
        method: "POST",
        body : JSON.stringify(_data),
        headers : {"Content-type" : "application/json; charset=UTD-8"}
    })
    .then(response => {
        if (response.status == 201) {
            //console.log("Logged In")
            window.open("login.html", "_self")
        }
    });
}
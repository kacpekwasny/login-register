

function setLoginAndToken(login, token) {
    if (localStorage.getItem("remember me") === "true") {
        localStorage.setItem("login", login)
        localStorage.setItem("token", token)
    } else {
        localStorage.removeItem("login")
        localStorage.removeItem("token")
    }
    setck("login", login)
    setck("token", token)
}

function removeLoginAndToken() {
    localStorage.removeItem("login")
    localStorage.removeItem("token")
    sessionStorage.removeItem("login")
    sessionStorage.removeItem("token")
    delck("login")
    delck("token")
}


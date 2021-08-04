function loadCredentials() {
    const login = localStorage.getItem("login")
    if (login===null) return
    const token = localStorage.getItem("token")
    if (token===null) return
    setck("login", login)
    setck("token", token)
}

loadCredentials()
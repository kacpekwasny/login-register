
////////////////////////////////////////////////////////////////
// Get needed elements once
const loginInp = document.getElementById('login-input-id');
const loginErr = document.getElementById("login-err-id");
const passInp = document.getElementById('password-input-id');
const passErr = document.getElementById("pass-err-id");
const loginBtn = document.getElementById("login-button-id");

// const checkMark = document.getElementById("check-mark-id");

////////////////////////////////////////////////////////////////
// Add event listeners

// login on ENTER-keydown
document.addEventListener("keydown", function (event) {
    if (event.keyCode === 13 || event.key === 13 || event.keyIdentifier === 13) {
        event.preventDefault();
        loginBtn.click();
    }
});


// to activate the Login button 
loginInp.addEventListener("keyup", watchInputs);
passInp.addEventListener("keyup", watchInputs);

// to discard the fail panel
loginInp.addEventListener("focus", watchInputs);
passInp.addEventListener("focus", watchInputs);


function watchInputs() {
    const logval = loginInp.value;
    const pasval = passInp.value;

    loginErr.innerText = ""
    passErr.innerText = "";

    if (logval.length < 1) {
        displayInfo(loginErr, "Login  can't be empty", "Login nie może być pusty")
        loginEnable(false)
        return
    }
    if (pasval.length < 1) {
        displayInfo(passErr, "Password can't be empty", "Hasło nie może być puste")
        loginEnable(false)
        return
    }
    loginEnable(true)
}

function loginEnable(on) {
    if (on) {
        loginBtn.classList.remove("disabled")
        loginBtn.style.cursor = ""
    } else {
        loginBtn.classList.add("disabled")
        loginBtn.style.cursor = "not-allowed"
    }
}

function clickLoginBtn() {
    if (loginBtn.classList.contains("disabled")) {
        return
    }
    sendLogin(loginInp.value, passInp.value);
}

function sendLogin(login, password) {
    const data = JSON.stringify({ login: login, password: password });

    fetch("/login", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: data
    })
        .then(res => {
            if (res.status != 200) {
                displayInfo(passErr, "Internal error: " + res.status, "Błąd wewnętrzny" + res.status);
            }
            res.json().then(j => {
                var has, lacking, tmp;
                tmp = hasProperties(j, ["err_code", "more"]);
                has = tmp[0]
                lacking = tmp[1]
                if (!has) {
                    displayInfo(passErr, "Internal error", "Błąd wewnętrzny");
                    console.error("Lacking keys in response JSON:", lacking)
                    return
                }
                // Check if successful registered
                switch (j.err_code) {
                    case 0:
                        // error code for 'ok'
                        break
                    case 1:
                        displayInfo(passErr, "Internal error", "Błąd wewnętrzny")
                        return
                    case 2:
                        displayInfo(loginErr, "Login not found", "Login nie został znaleziony")
                        return
                    case 3:
                        displayInfo(passErr, "Password doesn't match", "Hasło nie pasuje")
                        return
                    default:
                        displayInfo(passErr, "Internal error", "Błąd wewnętrzny")
                        return

                }
                // user succesfuly logged in, show check mark
                setck("login", login)
                setck("token", j.more.token)
                var link = getck("redirect_to");
                if (link == "") {
                    link = "/";
                } else {
                    delck("redirect_to");
                }
                window.location.href = link;
            });
        })
        .catch(err => {
            console.log("Cought an error!");
            console.error(err);
            displayError("Error: " + err, true);
        });
}

// if user was redirected from /register he will have 'new_account=true' in cookies.
// display blue info panel
function start() {
    var url = new URL(window.location.href)
    var newac = url.searchParams.get("new_account")
    if (newac === "true") {
        window.history.pushState({}, document.title, "/login")
        const s = document.getElementById("success-div-id")
        s.classList.remove("d-none")
        setTimeout(() => {
            s.classList.add("d-none")
            loginInp.focus()
        }, 3000)
    }
}



////////////////////////
// ————— MAIN ––––––– //
start();



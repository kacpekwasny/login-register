
////////////////////////////////////////////////////////////////
// Get needed elements once
const loginInp = document.getElementById('login-input-id');
const loginErr = document.getElementById('login-err-id');
const passInp = document.getElementById('password-input-id');
const passInp2 = document.getElementById('password-input-2-id');
const passErr = document.getElementById('password-err-id');
const registerBtn = document.getElementById("register-button-id");

// const checkMark = document.getElementById("check-mark-id");

////////////////////////////////////////////////////////////////
// Add event listeners

// register on keydown
document.addEventListener("keydown", function (event) {
    if (event.keyCode === 13 || event.key === 13 || event.keyIdentifier === 13) {
        event.preventDefault();
        registerBtn.click();
    }
});


// to activate the Register button 
loginInp.addEventListener("keyup", watchInputs);
passInp.addEventListener("keyup", watchInputs);
passInp2.addEventListener("keyup", watchInputs);

// to discard the fail panel
loginInp.addEventListener("focus", watchInputs);
passInp.addEventListener("focus", watchInputs);
passInp2.addEventListener("focus", watchInputs);


function watchInputs() {
    const logval = loginInp.value;
    const pasval = passInp.value;
    const pasval2 = passInp2.value;
    loginErr.innerText = "";
    passErr.innerText = "";
    if (logval.length < 2) {
        displayInfo(loginErr, "Login has to be at least 2 charachters", "Login musi mieć co najmniej 2 znaki")
        registerEnable(false)
        return
    }
    if (pasval.length < 7) {
        displayInfo(passErr, "Password cannot be shorter than 7 charachters", "Hasło nie może być krótsze niż 7 znaków")
        registerEnable(false)
        return
    }
    if (pasval !== pasval2) {
        displayInfo(passErr, "Passwords don't match", "Hasła się różnią");
        registerEnable(false)
        // return
    }

    passwordPwned(pasval).then(pwned => {
        if (pwned) {
            displayInfo(passErr, "Chose a different password, this one is popular", "Wybierz inne hasło, to jest popularne")
            registerEnable(false)
        return
        }
        if (pasval !== pasval2) registerEnable(true)

    }).catch(e => {
        console.error(e)
        registerEnable(true)
    })


}

function registerEnable(on) {
    if (on) {
        registerBtn.classList.remove("disabled")
        registerBtn.style.cursor = ""
    } else {
        registerBtn.classList.add("disabled")
        registerBtn.style.cursor = "not-allowed"
    }
}

function clickRegisterBtn() {
    console.log("clicked register btn")
    if (registerBtn.classList.contains("w3-disabled")) {
        return
    }
    console.log("  sending request")
    sendRegister(loginInp.value, passInp.value);
}

function sendRegister(login, password) {
    const data = JSON.stringify({ login: login, password: password });

    fetch("/register", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: data
    })
        .then(res => {
            if (res.status != 200) {
                displayInfo(passErr, "Internal error: " + res.status, "Błąd wewnętrzny" + res.status);
                console.log(res);
            }
            res.json().then(j => {
                var has, lacking, tmp;
                tmp = hasProperties(j, ["err_code"]);
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
                        displayInfo(passErr, "Internal error", "Błąd wewnętrzny");
                        return
                    case 5:
                        displayInfo(loginErr, "Login is in use", "Login w użyciu");
                        return
                    case 8:
                        displayInfo(passErr, "Chose a different password, this one is popular", "Wybierz inne hasło, to jest popularne")
                    default:
                        displayInfo(passErr, "Internal error", "Błąd wewnętrzny");
                        return

                }
                // user succesfuly registered, show check mark
                // show(checkMark);
                window.location.href = "/login?new_account=true"
            });
        })
        .catch(err => {
            console.log("Cought an error!");
            console.error(err);
            displayInfo(passErr, "Internal error", "Błąd wewnętrzny");
        });
}


function start() {
    loginInp.focus()
}


start()
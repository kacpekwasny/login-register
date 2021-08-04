
////////////////////////////////////////////////////////////////
// Get needed elements once
const loginInp = document.getElementById('login-input-id');
const loginErr = document.getElementById('login-err-id');
const passInp = document.getElementById('password-input-id');
const passInp2 = document.getElementById('password-input-2-id');
const passErr = document.getElementById('password-err-id');
const registerBtn = document.getElementById("register-button-id");

// check password for pwned only once
var password_not_pwned = " oa89 hdaiphu cp9<a8wh (P*j hapw8h a98wd pahsd pw98a dhpAIHJD pwa89"

// const checkMark = document.getElementById("check-mark-id");

////////////////////////////////////////////////////////////////
// Add event listeners

// register on keydown
document.addEventListener("keyup", function (event) {
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
    const invalid = loginValidChars(logval)
    if (invalid.length>0) {
        displayInfo(loginErr, "Do not use these characters in login: "+invalid, "Nie używaj tych znaków w loginie: "+invalid)
        return
    }
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
    if (password_not_pwned===pasval) {
        if (pasval === pasval2) registerEnable(true)
        return
    }
    passwordPwned(pasval).then(pwned => {
        if (pwned) {
            displayInfo(passErr, "Chose a different password, this one is popular", "Wybierz inne hasło, to jest popularne")
            registerEnable(false)
        return
        }
        password_not_pwned = pasval
        if (pasval === pasval2) registerEnable(true)

    }).catch(e => {
        console.error(e)
        if (pasval === pasval2) registerEnable(true)
    })


}

function loginValidChars(login) {
    const valid = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_ "
	var invalid = ""
	for (let ch of login) {
		if (!valid.includes(ch)) {
			invalid += `'${ch}', `
		}
	}
	return invalid
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
    if (registerBtn.classList.contains("disabled")) {
        return
    }
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
                        return
                    case 9:
                        displayInfo(loginErr, "Do not use these characters in login: "+j.more.invalid_chars, "Nie używaj tych znaków w loginie: "+j.more.invalid_chars)
                        return
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
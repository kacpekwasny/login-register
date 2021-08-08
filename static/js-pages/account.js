const newLoginInp = document.getElementById("new-login-input-id")
const newPaswdInp = document.getElementById("new-password-input-id")
const newPaswd2Inp = document.getElementById("new-password-2-input-id")
const paswdInp = document.getElementById("password-input-id")

const newLoginErr = document.getElementById("new-login-err-id")
const newPaswdErr = document.getElementById("new-password-err-id")
const paswdErr = document.getElementById("password-err-id")

const commitBtn = document.getElementById("commit-change-btn-id")
const logoutBtn = document.getElementById("log-out-btn-id")

// check password for pwned only once
var password_not_pwned = " oa89 hdaiphu cp9<a8wh (P*j hapw8h a98wd pahsd pw98a dhpAIHJD pwa89"

document.addEventListener("keyup", function (event) {
    if (event.keyCode === 13 || event.key === 13 || event.keyIdentifier === 13) {
        event.preventDefault();
        commitBtn.click();
    }
});

for (let elem of [newLoginInp, newPaswdInp, newPaswd2Inp, paswdInp]) {
    elem.addEventListener("keyup", watchInputs)
    elem.addEventListener("focus", watchInputs)
}


function watchInputs() {
    const logval = newLoginInp.value
    const npasval = newPaswdInp.value
    const npas2val = newPaswd2Inp.value
    const pasval = paswdInp.value
    newLoginErr.innerText = ""
    newPaswdErr.innerText = ""
    // watch new login
    if (logval.length === 0) {
        newLoginErr.classList.remove("text-danger")
        newLoginErr.classList.remove("text-warning")
        newLoginErr.classList.add("text-success")
        displayInfo(newLoginErr, "Login will not be changed", "Login nie zostanie zmieniony")
    } else {
        newLoginErr.classList.remove("text-success")
        newLoginErr.classList.add("text-danger")
        if (logval.length < 2) {
            displayInfo(newLoginErr, "Login has to be at least 2 charachters", "Login musi mieć co najmniej 2 znaki")
            commitEnable(false)
            return
        }
        var invalid = loginValidChars(logval)
        if (invalid.length > 0) {
            displayInfo(newLoginErr, "Do not use these characters in login: " + invalid, "Nie używaj tych znaków w loginie: " + invalid)
            commitEnable(false)
            return
        }
        newLoginErr.classList.remove("text-danger")
        newLoginErr.classList.add("text-warning")
        displayInfo(newLoginErr, "Login will be changed", "Login zostanie zmieniony")
        if (pasval.length > 0) commitEnable(true)
    }

    // watch new password
    if (npasval.length == 0) {
        newPaswdErr.classList.remove("text-warning")
        newPaswdErr.classList.remove("text-danger")
        newPaswdErr.classList.add("text-success")
        displayInfo(newPaswdErr, "Password will not be changed", "Hasło nie zostanie zmienione")
    } else {
        newPaswdErr.classList.remove("text-success")
        newPaswdErr.classList.add("text-danger")
        if (npasval.length < 7) {
            displayInfo(newPaswdErr, "Password cannot be shorter than 7 charachters", "Hasło nie może być krótsze niż 7 znaków")
            commitEnable(false)
            return
        }
        if (npasval !== npas2val) {
            displayInfo(newPaswdErr, "Passwords don't match", "Hasła się różnią");
            commitEnable(false)
            // return
        }
        if (password_not_pwned === npasval) {
            if (npasval === npas2val && pasval.length > 0) {
                newPaswdErr.classList.remove("text-danger")
                newPaswdErr.classList.add("text-warning")
                displayInfo(newPaswdErr, "Password will be changed", "Hasło zostanie zmienione")
                commitEnable(true)
            } else {
                commitEnable(false)
            }
            return
        }
        passwordPwned(npasval).then(pwned => {
            if (pwned) {
                displayInfo(newPaswdErr, "Chose a different password, this one is popular", "Wybierz inne hasło, to jest popularne")
                commitEnable(false)
                return
            }
            password_not_pwned = npasval
            if (npasval === npas2val && pasval.length>0) commitEnable(true)

        }).catch(e => {
            console.error(e)
            if (npasval === npas2val && pasval.length>0) commitEnable(true)
        })
    }
}

function clickCommitBtn() {
    if (commitBtn.classList.contains("disabled")) {
        return
    }
    sendChange(newLoginInp.value, newPaswdInp.value);
}

function clickLogoutBtn() {
    removeLoginAndToken()
    window.location.href = "/login"
}

function sendChange() {
    var login = getck("login")
    if (login==="") {
        login = localStorage.getItem("login")
        if (login==="") {
            login = sessionStorage.getItem("login")
            if (login==="") {
                console.error(`login not in cookie, local nor session storage, login===''`)
                alert("error, login=''")
                return
            }
        }
    }
    const data = JSON.stringify(
        {
            action: "change",
            login: login,
            password: paswdInp.value,
            newLogin: newLoginInp.value,
            newPassword: newPaswdInp.value,
        })
        console.log(data)
    fetch("/account", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: data
    })
        .then(res => {
            if (res.status != 200) {
                displayInfo(paswdErr, "Internal error: " + res.status, "Błąd wewnętrzny" + res.status);
            }
            res.json().then(j => {
                console.log(j)
                var has, lacking, tmp;
                tmp = hasProperties(j, ["err_code"]);
                has = tmp[0]
                lacking = tmp[1]
                if (!has) {
                    displayInfo(paswdErr, "Internal error", "Błąd wewnętrzny");
                    console.error("Lacking keys in response JSON:", lacking)
                    return
                }
                // Check if successful registered
                switch (j.err_code) {
                    case 0:
                        // error code for 'ok'
                        alertLang("Success!\nNow log in with new credentials.", "Sukces!\nTeraz zaloguj się przy użyciu nowych danych.")
                        clickLogoutBtn()
                        break
                    case 1:
                        displayInfo(paswdErr, "Internal error", "Błąd wewnętrzny");
                        return
                    case 3:
                        displayInfo(paswdErr, "Password doesn't match", "Hasło nie pasuje")
                        return
                    case 5:
                        displayInfo(newLoginErr, "Login is in use", "Login w użyciu");
                        return
                    case 8:
                        displayInfo(newPaswdErr, "Chose a different password, this one is popular", "Wybierz inne hasło, to jest popularne")
                        return
                    case 9:
                        displayInfo(newLoginErr, "Do not use these characters in login: " + j.more.invalid_chars, "Nie używaj tych znaków w loginie: " + j.more.invalid_chars)
                        return
                    default:
                        displayInfo(paswdErr, "Internal error", "Błąd wewnętrzny");
                        return

                }
            });
        })
        .catch(err => {
            console.log("Cought an error!");
            console.error(err);
            displayInfo(passErr, "Internal error", "Błąd wewnętrzny");
        });

}

function commitEnable(on) {
    if (on) {
        commitBtn.classList.remove("disabled")
        commitBtn.style.cursor = ""
    } else {
        commitBtn.classList.add("disabled")
        commitBtn.style.cursor = "not-allowed"
    }
}

newLoginInp.placeholder = getck("login")
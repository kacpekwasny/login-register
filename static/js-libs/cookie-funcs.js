
function setck(key, val) {
    document.cookie = `${key}=${val}; path=/`;
}

function delck(key) {
    document.cookie = `${key}= ; expires = Thu, 01 Jan 1970 00:00:00 GMT; path=/`
}

function getck(wanted_key) {
    const split = document.cookie.split(";");
    if (split[0]=="") {
        return ""
    }
    for (var i in split) {
        var key, val, tmp;
        tmp = split[i].split("=");
        key = tmp[0].trim();
        val = tmp[1].trim();
        if (key===wanted_key) {
            return val;
        }
    }
    return ""
}


function textByLang(en, pl) {
    const lang = getck("lang")
    switch (lang) {
        case "en":
            return en
        case "pl":
            return pl
        default:
            return en
    }
}
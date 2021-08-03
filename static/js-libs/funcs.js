// return true if an object has all specified properties and 'property_name' if it lacks that field
function hasProperties(obj, properties) {
    var lacking_props = [];
    for (let p in properties) {
        p = properties[p];
        if (!obj.hasOwnProperty(p)) {
            lacking_props.push(p);
        }       
    }
    if (lacking_props.length) {
        return [false, lacking_props]
    }
    return [true, undefined]
}


function removeParam(key, sourceURL) {
    var rtn = sourceURL.split("?")[0],
        param,
        params_arr = [],
        queryString = (sourceURL.indexOf("?") !== -1) ? sourceURL.split("?")[1] : "";
    if (queryString !== "") {
        params_arr = queryString.split("&");
        for (var i = params_arr.length - 1; i >= 0; i -= 1) {
            param = params_arr[i].split("=")[0];
            if (param === key) {
                params_arr.splice(i, 1);
            }
        }
        if (params_arr.length) rtn = rtn + "?" + params_arr.join("&");
    }
    return rtn;
}


function displayPassword(id) {
    var pass = document.getElementById(id)
    if (pass.type==="password") {
        pass.type = "text"
    } else {
        pass.type = "password"
    }
}


// return true if an object has all specified properties and 'property_name' if it lacks that field
function hasProperties(obj, properties) {
    var lacking_props = [];
    for (let p in properties) {
        p = properties[p];
        if (!obj.hasOwnProperty(p)) {
            lacking_props.push(p);
        }       
    }
    if (lacking_props.length) {
        return [false, lacking_props]
    }
    return [true, undefined]
}


function displayInfo(elem, en, pl) {
    elem.innerText = textByLang(en, pl)
}

function alertLang(en, pl) {
    alert(textByLang(en, pl))
}
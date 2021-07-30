

async function getSHA1(message) {
    const msgUint8 = new TextEncoder().encode(message);                           // encode as (utf-8) Uint8Array
    const hashBuffer = await crypto.subtle.digest('SHA-1', msgUint8);           // hash the message
    const hashArray = Array.from(new Uint8Array(hashBuffer));                     // convert buffer to byte array
    const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join(''); // convert bytes to hex string
    return hashHex;
  }

  async function getSHA256(message) {
    const msgUint8 = new TextEncoder().encode(message);                           // encode as (utf-8) Uint8Array
    const hashBuffer = await crypto.subtle.digest('SHA-256', msgUint8);           // hash the message
    const hashArray = Array.from(new Uint8Array(hashBuffer));                     // convert buffer to byte array
    const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join(''); // convert bytes to hex string
    return hashHex;
  }

async function passwordPwned(paswd) {
    // return true if password was leaked and is in a database
    const hash = await getSHA1(paswd);
    const hash5 = hash.slice(0, 5);
    const hashend = hash.slice(-35).toUpperCase();
    const res = await fetch("https://api.pwnedpasswords.com/range/"+hash5,{method: "GET"});
    var text = await res.text();
    text = text.replace("\r", "");
    var lines = text.split("\n");
    for (var i=0; i<lines.length; i++) {
        lines[i] = lines[i].split(":")[0];
        if (lines[i]===hashend) {
            return true;
        }       
    }
    return false;
}



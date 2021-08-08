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

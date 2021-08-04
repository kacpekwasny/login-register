package main

// var Messages = map[string]string{
// 	"internal_error": "Internal error",
//
// 	// login
// 	"login_not_found": "Login was not found",
// 	"pass_missmatch":  "Password doesn't match",
//
// 	// register
// 	"login_in_use":       "Login is in use",
// 	"login_requirements": "Login doesn't match requirements",
// 	"pass_requirements":  "Password doesn't match requirements",
// 	"paswd_pwned":        "Password has been leaked and can be easily hacked",
// }

var ErrCodes = map[string]int{
	"ok":             0,
	"internal_error": 1,

	"login_not_found": 2,
	"pass_missmatch":  3,

	"login_in_use":       5,
	"login_requirements": 6,
	"pass_requirements":  7,
	"paswd_pwned":        8,
	"invalid_chars":      9,
}

func GetErrCode(message_title string) int {
	code, ok := ErrCodes[message_title]
	if !ok {
		M.Log1("ErrCodes[ %v ] is missing", message_title)
		return 1
	}
	return code
}

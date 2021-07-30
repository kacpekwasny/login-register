package main

import (
	"fmt"
)

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
}

//func GetMessage(message_title string, r *http.Request) string {
//	if message_title == "ok" {
//		return "ok"
//	}
//	key := message_title + "_" + GetLangCookie(r)
//	m, ok := Messages[key]
//	if !ok {
//		fmt.Println("Messages[" + key + "] is missing")
//		m, ok := Messages[message_title]
//		if !ok {
//			fmt.Println("Messages[" + message_title + "] is missing")
//			return "Message is missing!"
//		}
//		return m
//	}
//	return m
//}

func GetErrCode(message_title string) int {
	code, ok := ErrCodes[message_title]
	if !ok {
		fmt.Println("ErrCodes[ " + message_title + " ] is missing")
		return 1
	}
	return code
}

package main

const (
	cookie_info_en = "Using this website is equal to accepting the use of cookies."
	cookie_info_pl = "Używanie tej strony jest równoznaczne z akceptacją używania ciasteczek."

	nav_go_home_en = "Home"
	nav_go_home_pl = "Główna"

	nav_log_in_en = "Log In"
	nav_log_in_pl = "Zaloguj się"

	nav_new_acc_en = "Sign up"
	nav_new_acc_pl = "Nowe konto"

	action_log_in_en = "Log in"
	action_log_in_pl = "Zaloguj się"

	action_new_acc_en = "Create account"
	action_new_acc_pl = "Utwórz konto"
)

var page_map = map[string]map[string]string{
	"login_en":    login_en,
	"login_pl":    login_pl,
	"register_en": register_en,
	"register_pl": register_pl,
	"account_en":  account_en,
	"account_pl":  account_pl,
}

var login_en = map[string]string{
	"title":           nav_log_in_en,
	"nav go home":     nav_go_home_en,
	"nav log in":      nav_log_in_en,
	"nav register":    nav_new_acc_en,
	"password input":  "Password",
	"log in action":   action_log_in_en,
	"cookie use info": cookie_info_en,
	"remember me":     "Remember me",
}

var login_pl = map[string]string{
	"title":           nav_log_in_pl,
	"nav go home":     nav_go_home_pl,
	"nav log in":      nav_log_in_pl,
	"nav register":    nav_new_acc_pl,
	"password input":  "Hasło",
	"log in action":   action_log_in_pl,
	"cookie use info": cookie_info_pl,
	"remember me":     "Zapamiętaj mnie",
}

var register_en = map[string]string{
	"title":                  nav_new_acc_en,
	"nav go home":            nav_go_home_en,
	"nav log in":             nav_log_in_en,
	"nav register":           nav_new_acc_en,
	"password input":         "Password",
	"repeat password input":  "Repeat it",
	"register action button": action_new_acc_en,
	"cookie use info":        cookie_info_en,
}

var register_pl = map[string]string{
	"title":                  "Rejestracja",
	"nav go home":            nav_go_home_pl,
	"nav log in":             nav_log_in_pl,
	"nav register":           nav_new_acc_pl,
	"password input":         "Hasło",
	"repeat password input":  "Powtórz",
	"register action button": action_new_acc_pl,
	"cookie use info":        cookie_info_pl,
}

var account_en = map[string]string{
	"title":                     "Account",
	"nav go home":               nav_go_home_en,
	"nav log out":               "Log out",
	"account security settings": "Change login and password",
	"new password":              "New password",
	"repeat new password":       "Repeat it",
	"current password":          "Current password",
	"commit changes btn":        "Confirm",
	"cookie use info":           cookie_info_en,
}

var account_pl = map[string]string{
	"title":                     "Konto",
	"nav go home":               nav_go_home_pl,
	"nav log out":               "Wyloguj",
	"account security settings": "Zmień login i hasło",
	"new password":              "Nowe hasło",
	"repeat new password":       "Powtórz je",
	"current password":          "Obecne hasło",
	"commit changes btn":        "Zatwierdź",
	"cookie use info":           cookie_info_pl,
}

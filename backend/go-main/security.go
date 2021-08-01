package main

//var (
//	hashKey  = []byte("HnwKnhBdZIHduY!1af")
//	blockKey = []byte("^rKYpckU7f42MD^w")
//)
//var cookieSec = securecookie.New(hashKey, blockKey)
//
//func setCookie(w http.ResponseWriter, cookie_name string, key_val_pairs ...string) {
//	l := len(key_val_pairs)
//	if l%2 != 0 {
//		panic(ErrInvalidInput)
//	}
//	value := map[string]string{}
//	for i := 0; i < l-2; i += 2 {
//		value[key_val_pairs[i]] = key_val_pairs[i+1]
//	}
//	encoded_val, err := cookieSec.Encode("cookies", value)
//	cmt.Pc(err)
//	cookie := &http.Cookie{
//		Name:     cookie_name,
//		Value:    encoded_val,
//		Path:     "/",
//		Secure:   true,
//		HttpOnly: true,
//	}
//	http.SetCookie(w, cookie)
//}
//

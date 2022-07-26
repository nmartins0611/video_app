package helpers

import "net/http"

//
// SetCookie : Set a cookie's value
func SetCookie(w http.ResponseWriter, name string, value string, age int) {
	http.SetCookie(w, &http.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: age,
	})
}

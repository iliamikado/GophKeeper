package handlers

import (
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	login, err := passwordManager.Register(user.Login, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token := buildJWTString(login)
	http.SetCookie(w, &http.Cookie{Name: "JWT", Value: token})
}
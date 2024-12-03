package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	login, err := passwordManager.Login(user.Login, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token := buildJWTString(login)
	http.SetCookie(w, &http.Cookie{Name: "JWT", Value: token})
}

func buildJWTString(login string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{},
		Login: login,
	})

	tokenString, _ := token.SignedString([]byte(SecretKey))
	return tokenString
}
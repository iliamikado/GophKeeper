package handlers

import (
	"PasswordManager/internal/logger"
	"context"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

// SecretKey - секретный ключ
const SecretKey = "secret key"

type loginKey struct{}

// JWTClaims - данные в jwt
type JWTClaims struct {
	jwt.RegisteredClaims
	Login string
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("JWT")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := cookie.Value
		login, err := getLogin(token)
		if err != nil {
			logger.Info(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !passwordManager.CheckLogin(login) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), loginKey{}, login))
		next.ServeHTTP(w, r)
	}
}

func getLogin(tokenString string) (string, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("token is not valid")
	}

	return claims.Login, nil
}

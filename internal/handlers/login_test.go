package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	initApp()
	user := User{
		Login: "aaa",
		Password: "bbb",
	}
	passwordManager.Register(user.Login, user.Password)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", dataToBody(user))
	rr := httptest.NewRecorder()

	Login(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	header := rr.Header().Get("Set-Cookie")
	assert.NotEmpty(t, header)

	req = httptest.NewRequest(http.MethodPost, "/api/v1/login", dataToBody(User{}))
	rr = httptest.NewRecorder()

	Login(rr, req)

	assert.Equal(t, rr.Code, http.StatusUnauthorized)
}
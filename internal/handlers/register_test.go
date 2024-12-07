package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	initApp()
	user := User{
		Login: "aaa",
		Password: "bbb",
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register", dataToBody(user))
	rr := httptest.NewRecorder()

	Register(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	header := rr.Header().Get("Set-Cookie")
	assert.NotEmpty(t, header)

	req = httptest.NewRequest(http.MethodPost, "/api/v1/register", dataToBody(user))
	rr = httptest.NewRecorder()

	Register(rr, req)

	assert.Equal(t, rr.Code, http.StatusUnauthorized)
}
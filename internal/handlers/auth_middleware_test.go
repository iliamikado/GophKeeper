package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	initApp()

	req := httptest.NewRequest(http.MethodPost, "/get_all", nil)
	rr := httptest.NewRecorder()

	authMiddleware(GetAllData)(rr, req)

	assert.Equal(t, rr.Code, http.StatusUnauthorized)
	
	passwordManager.Register(defaultLogin, "a")
	token := buildJWTString(defaultLogin)

	req = httptest.NewRequest(http.MethodPost, "/get_all", nil)
	req.AddCookie(&http.Cookie{
		Name: "JWT",
		Value: token,
	})
	rr = httptest.NewRecorder()

	authMiddleware(GetAllData)(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}
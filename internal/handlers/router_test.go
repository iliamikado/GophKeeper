package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppRouter(t *testing.T) {
	initApp()
	r := AppRouter()

	user := User{
		Login: "aaa",
		Password: "bbb",
	}

	req := httptest.NewRequest(http.MethodPost, "/register", dataToBody(user))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	req = httptest.NewRequest(http.MethodPost, "/login", dataToBody(user))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}
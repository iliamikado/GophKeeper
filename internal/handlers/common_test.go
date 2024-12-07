package handlers

import (
	"PasswordManager/internal/manager"
	"bytes"
	"encoding/json"
	"io"
)

const defaultLogin string = "login"

func initApp() {
	manager.Initialize()
	passwordManager = manager.AppManager
	passwordManager.Register(defaultLogin, "a")
}

func dataToBody(data any) io.Reader {
	b, _ := json.Marshal(data)
	return bytes.NewReader(b)
}
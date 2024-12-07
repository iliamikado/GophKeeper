// Пакет с придожением
package app

import (
	"PasswordManager/internal/handlers"
	"PasswordManager/internal/logger"
	"PasswordManager/internal/manager"
	"net/http"
)

// StartServer - запуск сервера
func StartServer() {
	manager.Initialize()
	r := handlers.AppRouter()

	logger.Info("Running server on port " + ":8080")
	err := http.ListenAndServe(":8080", r)
	if (err != nil) {
		logger.Panic(err)
	}
}
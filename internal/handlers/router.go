// Пакет с ручками
package handlers

import (
	"PasswordManager/internal/manager"

	"github.com/go-chi/chi/v5"
)

var passwordManager *manager.IPasswordManager

func AppRouter() *chi.Mux {
	passwordManager = manager.PasswordManager
	r := chi.NewRouter()
	r.Post("/register", Register)
	r.Post("/login", Login)
	r.Post("/save_enter_data", authMiddleware(SaveEnterData))
	r.Get("/get_enter_data", authMiddleware(GetEnterData))
	r.Post("/save_text_data", authMiddleware(SaveTextData))
	r.Get("/get_text_data", authMiddleware(GetTextData))
	r.Post("/save_card_data", authMiddleware(SaveBankCardData))
	r.Get("/get_card_data", authMiddleware(GetBankCardData))
	r.Get("/get_all_data", authMiddleware(GetAllData))
	return r
}
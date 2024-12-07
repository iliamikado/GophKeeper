// Пакет с ручками
package handlers

import (
	"PasswordManager/internal/manager"

	"github.com/go-chi/chi/v5"
)

var passwordManager *manager.PasswordManager

func AppRouter() *chi.Mux {
	passwordManager = manager.AppManager
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/register", Register)
		r.Post("/login", Login)
		r.Route("/enter_data", func(r chi.Router) {
			r.Post("/save", authMiddleware(SaveEnterData))
			r.Get("/get", authMiddleware(GetEnterData))
		})
		r.Route("/text_data", func(r chi.Router) {
			r.Post("/save", authMiddleware(SaveTextData))
			r.Get("/get", authMiddleware(GetTextData))
		})
		r.Route("/payment_card", func(r chi.Router) {
			r.Post("/save", authMiddleware(SavePaymentCard))
			r.Get("/get", authMiddleware(GetPaymentCard))
		})

		r.Get("/get_all", authMiddleware(GetAllData))
	})
	return r
}

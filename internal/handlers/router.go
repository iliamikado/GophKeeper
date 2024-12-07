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
			r.Post("/", authMiddleware(SaveEnterData))
			r.Get("/", authMiddleware(GetEnterData))
		})
		r.Route("/text_data", func(r chi.Router) {
			r.Post("/", authMiddleware(SaveTextData))
			r.Get("/", authMiddleware(GetTextData))
		})
		r.Route("/payment_card", func(r chi.Router) {
			r.Post("/", authMiddleware(SavePaymentCard))
			r.Get("/", authMiddleware(GetPaymentCard))
		})

		r.Get("/get_all", authMiddleware(GetAllData))
	})
	return r
}

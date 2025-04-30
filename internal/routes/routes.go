package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"qrcode-backend/internal/handlers"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/generate-qr", handlers.GenerateQRCode)
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return r
}
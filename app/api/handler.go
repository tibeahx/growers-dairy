package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tibeahx/growers-dairy/app/usecase"
)

type Handler struct {
	service *usecase.ServiceProvider
}

func NewHandler(service *usecase.ServiceProvider) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewMux()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/strains", func(r chi.Router) {
		r.Post("/create", http.HandlerFunc(h.CreateStrain))
		r.Post("/update", http.HandlerFunc(h.UpdateStrain))
		r.Get("/{id}", http.HandlerFunc(h.StrainByID))
		r.Delete("/{id}", http.HandlerFunc(h.DeleteStrainByID))
		r.Get("/list", http.HandlerFunc(h.Strains))
	})

	router.Route("/growlogs", func(r chi.Router) {
		r.Post("/create", http.HandlerFunc(h.CreateNewGrowLog))
		r.Post("/update", http.HandlerFunc(h.UpdateGrowLog))
		r.Get("/{id}", http.HandlerFunc(h.GrowLogByID))
		r.Delete("/{id}", http.HandlerFunc(h.DeleteGrowLogByID))
		r.Get("/list", http.HandlerFunc(h.GrowLogs))
	})

	router.Route("/logentries", func(r chi.Router) {
		r.Post("/create", http.HandlerFunc(h.CreateLogEntry))
		r.Post("/update", http.HandlerFunc(h.UpdateEntry))
		r.Get("/{id}", http.HandlerFunc(h.LogEntryByID))
		r.Delete("/{id}", http.HandlerFunc(h.DeleteEntryByID))
		r.Get("/list", http.HandlerFunc(h.LogEntries))
		r.Post(("/photos/upload"), http.HandlerFunc(h.UploadToBucket))
	})
	return router
}

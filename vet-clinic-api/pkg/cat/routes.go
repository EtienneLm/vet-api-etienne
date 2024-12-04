package cat

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	catConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/", catConfig.CreateCatHandler)
	router.Get("/", catConfig.GetCatsHandler)
	router.Get("/{id}", catConfig.GetCatByIDHandler)
	router.Put("/{id}", catConfig.UpdateCatHandler)
	router.Delete("/{id}", catConfig.DeleteCatHandler)

	return router
}

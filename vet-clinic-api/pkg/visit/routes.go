package visit

import (
	"net/http"
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) http.Handler {
	visitController := New(configuration)
	router := chi.NewRouter()

	router.Post("/", visitController.CreateVisitHandler)
	router.Get("/", visitController.GetVisitsHandler)
	router.Get("/{id}", visitController.GetVisitByIDHandler)
	router.Put("/{id}", visitController.UpdateVisitHandler)
	router.Delete("/{id}", visitController.DeleteVisitHandler)

	return router
}

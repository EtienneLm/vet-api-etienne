package treatment

import (
	"net/http"
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) http.Handler {
	treatmentController := New(configuration)
	router := chi.NewRouter()

	router.Post("/", treatmentController.CreateTreatmentHandler)
	router.Get("/", treatmentController.GetTreatmentsHandler)
	router.Get("/{id}", treatmentController.GetTreatmentByIDHandler)
	router.Put("/{id}", treatmentController.UpdateTreatmentHandler)
	router.Delete("/{id}", treatmentController.DeleteTreatmentHandler)

	return router
}

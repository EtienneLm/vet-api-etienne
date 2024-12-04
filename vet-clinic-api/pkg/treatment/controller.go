package treatment

import (
	"net/http"
	"strconv"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type TreatmentController struct {
	*config.Config
}

func New(config *config.Config) *TreatmentController {
	return &TreatmentController{Config: config}
}

func (controller *TreatmentController) CreateTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.TreatmentRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	treatment := &dbmodel.Treatment{
		VisitID:   req.VisitID,
		Name:      req.Name,
		Dosage:    req.Dosage,
		Frequency: req.Frequency,
	}

	newTreatment, err := controller.TreatmentRepository.Create(treatment)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create treatment"})
		return
	}

	render.JSON(w, r, newTreatment)
}

func (controller *TreatmentController) GetTreatmentsHandler(w http.ResponseWriter, r *http.Request) {
	treatments, err := controller.TreatmentRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve treatments"})
		return
	}

	render.JSON(w, r, treatments)
}

func (controller *TreatmentController) GetTreatmentByIDHandler(w http.ResponseWriter, r *http.Request) {
	treatmentID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(treatmentID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid treatment ID"})
		return
	}

	treatment, err := controller.TreatmentRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Treatment not found"})
		return
	}

	render.JSON(w, r, treatment)
}

func (controller *TreatmentController) UpdateTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.TreatmentRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	treatmentID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(treatmentID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid treatment ID"})
		return
	}

	treatment, err := controller.TreatmentRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Treatment not found"})
		return
	}

	treatment.Name = req.Name
	treatment.Dosage = req.Dosage
	treatment.Frequency = req.Frequency

	updatedTreatment, err := controller.TreatmentRepository.Update(treatment)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update treatment"})
		return
	}

	render.JSON(w, r, updatedTreatment)
}

func (controller *TreatmentController) DeleteTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	treatmentID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(treatmentID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid treatment ID"})
		return
	}

	if err := controller.TreatmentRepository.Delete(uint(id)); err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete treatment"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "Treatment deleted successfully"})
}

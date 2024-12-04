package visit

import (
	"net/http"
	"strconv"
	"time"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"

	"vet-clinic-api/pkg/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type VisitController struct {
	*config.Config
}

func New(config *config.Config) *VisitController {
	return &VisitController{Config: config}
}

func (controller *VisitController) CreateVisitHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.VisitRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	visit := &dbmodel.Visit{
		CatID:        req.CatID,
		Date:         time.Now(),
		Reason:       req.Reason,
		Veterinarian: req.Veterinarian,
	}

	newVisit, err := controller.VisitRepository.Create(visit)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create visit"})
		return
	}

	render.JSON(w, r, newVisit)
}

func (controller *VisitController) GetVisitsHandler(w http.ResponseWriter, r *http.Request) {
	visits, err := controller.VisitRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve visits"})
		return
	}

	render.JSON(w, r, visits)
}

func (controller *VisitController) GetVisitByIDHandler(w http.ResponseWriter, r *http.Request) {
	visitID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(visitID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid visit ID"})
		return
	}

	visit, err := controller.VisitRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Visit not found"})
		return
	}

	render.JSON(w, r, visit)
}

func (controller *VisitController) UpdateVisitHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.VisitRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	visitID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(visitID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid visit ID"})
		return
	}

	visit, err := controller.VisitRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Visit not found"})
		return
	}

	visit.Reason = req.Reason
	visit.Veterinarian = req.Veterinarian

	updatedVisit, err := controller.VisitRepository.Update(visit)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update visit"})
		return
	}

	render.JSON(w, r, updatedVisit)
}

func (controller *VisitController) DeleteVisitHandler(w http.ResponseWriter, r *http.Request) {
	visitID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(visitID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid visit ID"})
		return
	}

	if err := controller.VisitRepository.Delete(uint(id)); err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete visit"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "Visit deleted successfully"})
}

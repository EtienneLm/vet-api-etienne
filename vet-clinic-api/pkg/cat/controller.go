package cat

import (
	"net/http"
	"strconv"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type CatController struct {
	*config.Config
}

func New(config *config.Config) *CatController {
	return &CatController{Config: config}
}

func (controller *CatController) CreateCatHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	cat := &dbmodel.Cat{
		Name:   req.Name,
		Age:    req.Age,
		Breed:  req.Breed,
		Weight: req.Weight,
	}

	newCat, err := controller.CatRepository.Create(cat)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create cat"})
		return
	}

	render.JSON(w, r, newCat)
}

func (controller *CatController) GetCatsHandler(w http.ResponseWriter, r *http.Request) {
	cats, err := controller.CatRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve cats"})
		return
	}

	render.JSON(w, r, cats)
}

func (controller *CatController) GetCatByIDHandler(w http.ResponseWriter, r *http.Request) {
	catID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(catID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid cat ID"})
		return
	}

	cat, err := controller.CatRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Cat not found"})
		return
	}

	render.JSON(w, r, cat)
}

func (controller *CatController) UpdateCatHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	catID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(catID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid cat ID"})
		return
	}

	cat, err := controller.CatRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Cat not found"})
		return
	}

	cat.Name = req.Name
	cat.Age = req.Age
	cat.Breed = req.Breed
	cat.Weight = req.Weight

	updatedCat, err := controller.CatRepository.Update(cat)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update cat"})
		return
	}

	render.JSON(w, r, updatedCat)
}

func (controller *CatController) DeleteCatHandler(w http.ResponseWriter, r *http.Request) {
	catID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(catID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid cat ID"})
		return
	}

	if err := controller.CatRepository.Delete(uint(id)); err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete cat"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "Cat deleted successfully"})
}

package models

import (
	"net/http"

	"github.com/go-chi/render"
)

type TreatmentRequest struct {
	VisitID   uint   `json:"visit_id"`
	Name      string `json:"name"`
	Dosage    string `json:"dosage"`
	Frequency string `json:"frequency"`
}

func (tr *TreatmentRequest) Bind(r *http.Request) error {
	if err := render.Bind(r, tr); err != nil {
		return err
	}
	return nil
}

type TreatmentResponse struct {
	ID        uint   `json:"id"`
	VisitID   uint   `json:"visit_id"`
	Name      string `json:"name"`
	Dosage    string `json:"dosage"`
	Frequency string `json:"frequency"`
}

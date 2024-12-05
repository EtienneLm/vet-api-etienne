package models

import (
	"errors"
	"net/http"
)

type TreatmentRequest struct {
	VisitID   uint   `json:"visit_id"`
	Name      string `json:"name"`
	Dosage    string `json:"dosage"`
	Frequency string `json:"frequency"`
}

func (tr *TreatmentRequest) Bind(r *http.Request) error {
	if tr.VisitID <= 0 {
		return errors.New("id must be a positive number")
	}
	if tr.Name == "" {
		return errors.New("name field cannot be empty")
	}
	if tr.Dosage == "" {
		return errors.New("dosage field cannot be empty")
	}
	if tr.Frequency == "" {
		return errors.New("frequency must be a positive number")
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

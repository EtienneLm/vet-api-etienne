package models

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type VisitRequest struct {
	CatID        uint      `json:"cat_id"`
	Date         time.Time `json:"date"`
	Reason       string    `json:"reason"`
	Veterinarian string    `json:"veterinarian"`
}

func (vr *VisitRequest) Bind(r *http.Request) error {
	if err := render.Bind(r, vr); err != nil {
		return err
	}
	return nil
}

type VisitResponse struct {
	ID           uint      `json:"id"`
	CatID        uint      `json:"cat_id"`
	Date         time.Time `json:"date"`
	Reason       string    `json:"reason"`
	Veterinarian string    `json:"veterinarian"`
}

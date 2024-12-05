package models

import (
	"errors"
	"net/http"
	"time"
)

type VisitRequest struct {
	CatID  uint      `json:"cat_id"`
	Date   time.Time `json:"date"`
	Reason string    `json:"reason"`
	// Veterinarian string    `json:"veterinarian"`
}

func (vr *VisitRequest) Bind(r *http.Request) error {
	if vr.CatID <= 0 {
		return errors.New("cat_id must be a positive number")
	}
	if vr.Date.IsZero() {
		return errors.New("date field cannot be empty or invalid")
	}
	if vr.Reason == "" {
		return errors.New("reason field cannot be empty")
	}
	return nil
}

type VisitResponse struct {
	ID     uint      `json:"id"`
	CatID  uint      `json:"cat_id"`
	Date   time.Time `json:"date"`
	Reason string    `json:"reason"`
	// Veterinarian string    `json:"veterinarian"`
}

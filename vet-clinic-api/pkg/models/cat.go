package models

import (
	"errors"
	"net/http"
)

type CatRequest struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Breed  string  `json:"breed"`
	Weight float32 `json:"weight"`
}

type CatResponse struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Breed  string  `json:"breed"`
	Weight float32 `json:"weight"`
}

func (cr *CatRequest) Bind(r *http.Request) error {
	if cr.Name == "" {
		return errors.New("name field cannot be empty")
	}
	if cr.Age <= 0 {
		return errors.New("age must be a positive number")
	}
	if cr.Breed == "" {
		return errors.New("breed field cannot be empty")
	}
	if cr.Weight <= 0 {
		return errors.New("weight must be a positive number")
	}
	return nil
}

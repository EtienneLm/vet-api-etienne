package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Visit struct {
	gorm.Model
	CatID        uint      `json:"cat_id"`
	Date         time.Time `json:"date"`
	Reason       string    `json:"reason"`
	Veterinarian string    `json:"veterinarian"`
}

type VisitRepository interface {
	Create(visit *Visit) (*Visit, error)
	FindAll() ([]*Visit, error)
	FindByID(id uint) (*Visit, error)
	FindByCatID(catID uint) ([]*Visit, error)
	Update(visit *Visit) (*Visit, error)
	Delete(id uint) error
}

type visitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &visitRepository{db: db}
}

func (r *visitRepository) Create(visit *Visit) (*Visit, error) {
	if err := r.db.Create(visit).Error; err != nil {
		return nil, err
	}
	return visit, nil
}

func (r *visitRepository) FindAll() ([]*Visit, error) {
	var visits []*Visit
	if err := r.db.Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

func (r *visitRepository) FindByID(id uint) (*Visit, error) {
	var visit Visit
	if err := r.db.First(&visit, id).Error; err != nil {
		return nil, err
	}
	return &visit, nil
}

func (r *visitRepository) FindByCatID(catID uint) ([]*Visit, error) {
	var visits []*Visit
	if err := r.db.Where("cat_id = ?", catID).Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

func (r *visitRepository) Update(visit *Visit) (*Visit, error) {
	if err := r.db.Save(visit).Error; err != nil {
		return nil, err
	}
	return visit, nil
}

func (r *visitRepository) Delete(id uint) error {
	if err := r.db.Delete(&Visit{}, id).Error; err != nil {
		return err
	}
	return nil
}

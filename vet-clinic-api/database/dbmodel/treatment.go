package dbmodel

import (
	"gorm.io/gorm"
)

type Treatment struct {
	gorm.Model
	VisitID   uint   `json:"visit_id"`
	Name      string `json:"name"`
	Dosage    string `json:"dosage"`
	Frequency string `json:"frequency"`
}

type TreatmentRepository interface {
	Create(treatment *Treatment) (*Treatment, error)
	FindAll() ([]*Treatment, error)
	FindByID(id uint) (*Treatment, error)
	FindByVisitID(visitID uint) ([]*Treatment, error)
	Update(treatment *Treatment) (*Treatment, error)
	Delete(id uint) error
}

type treatmentRepository struct {
	db *gorm.DB
}

func NewTreatmentRepository(db *gorm.DB) TreatmentRepository {
	return &treatmentRepository{db: db}
}

func (r *treatmentRepository) Create(treatment *Treatment) (*Treatment, error) {
	if err := r.db.Create(treatment).Error; err != nil {
		return nil, err
	}
	return treatment, nil
}

func (r *treatmentRepository) FindAll() ([]*Treatment, error) {
	var treatments []*Treatment
	if err := r.db.Find(&treatments).Error; err != nil {
		return nil, err
	}
	return treatments, nil
}

func (r *treatmentRepository) FindByID(id uint) (*Treatment, error) {
	var treatment Treatment
	if err := r.db.First(&treatment, id).Error; err != nil {
		return nil, err
	}
	return &treatment, nil
}

func (r *treatmentRepository) FindByVisitID(visitID uint) ([]*Treatment, error) {
	var treatments []*Treatment
	if err := r.db.Where("visit_id = ?", visitID).Find(&treatments).Error; err != nil {
		return nil, err
	}
	return treatments, nil
}

func (r *treatmentRepository) Update(treatment *Treatment) (*Treatment, error) {
	if err := r.db.Save(treatment).Error; err != nil {
		return nil, err
	}
	return treatment, nil
}
func (r *treatmentRepository) Delete(id uint) error {
	if err := r.db.Delete(&Treatment{}, id).Error; err != nil {
		return err
	}
	return nil
}

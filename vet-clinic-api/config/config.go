package config

import (
	"vet-clinic-api/database"
	"vet-clinic-api/database/dbmodel"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DB                  *gorm.DB
	CatRepository       dbmodel.CatRepository
	VisitRepository     dbmodel.VisitRepository
	TreatmentRepository dbmodel.TreatmentRepository
}

func New() (*Config, error) {
	config := Config{}

	databaseSession, err := gorm.Open(sqlite.Open("vet_clinic.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)

	config.CatRepository = dbmodel.NewCatRepository(databaseSession)
	config.VisitRepository = dbmodel.NewVisitRepository(databaseSession)
	config.TreatmentRepository = dbmodel.NewTreatmentRepository(databaseSession)

	return &config, nil
}

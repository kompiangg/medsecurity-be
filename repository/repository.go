package repository

import (
	"medsecurity/config"
	"medsecurity/repository/access_history"
	"medsecurity/repository/cloudinary"
	"medsecurity/repository/doctor"
	"medsecurity/repository/patient"
	"medsecurity/repository/patient_image"
	"medsecurity/repository/patient_secret"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Doctor        doctor.Repository
	Patient       patient.Repository
	PatientSecret patient_secret.Repository
	PatientImage  patient_image.Repository
	Cloudinary    cloudinary.Repository
	AccessHistory access_history.Repository
}

func New(
	config config.Config,
	db *sqlx.DB,
	redis *redis.Client,
) (Repository, error) {
	cloudinary, err := cloudinary.New(cloudinary.Config{
		URIConnection: config.Cloudinary.URIConnection,
	})
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		Doctor:        doctor.New(db),
		Patient:       patient.New(db),
		PatientSecret: patient_secret.New(db),
		PatientImage:  patient_image.New(db, redis),
		AccessHistory: access_history.New(db),
		Cloudinary:    cloudinary,
	}, nil
}

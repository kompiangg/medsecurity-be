package repository

import (
	"medsecurity/config"
	"medsecurity/repository/cloudinary"
	"medsecurity/repository/doctor"
	"medsecurity/repository/patient"
	"medsecurity/repository/patient_image"
	"medsecurity/repository/patient_secret"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Doctor        doctor.Repository
	Patient       patient.Repository
	PatientSecret patient_secret.Repository
	PatientImage  patient_image.Repository
	Cloudinary    cloudinary.Repository
}

func New(
	config config.Config,
	db *sqlx.DB,
	// redis *redis.Client,
	// cld objstorage.ObjectStorageItf,
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
		PatientImage:  patient_image.New(db),
		Cloudinary:    cloudinary,
	}, nil
}

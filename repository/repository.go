package repository

import (
	"medsecurity/config"
	"medsecurity/repository/auth"
	"medsecurity/repository/doctor"
	"medsecurity/repository/patient"
	"medsecurity/repository/patient_secret"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Auth          auth.Repository
	Doctor        doctor.Repository
	Patient       patient.Repository
	PatientSecret patient_secret.Repository
}

func New(
	config config.Config,
	db *sqlx.DB,
	// redis *redis.Client,
	// cld objstorage.ObjectStorageItf,
) (Repository, error) {
	return Repository{
		Auth:          auth.New(db),
		Doctor:        doctor.New(db),
		Patient:       patient.New(db),
		PatientSecret: patient_secret.New(db),
	}, nil
}

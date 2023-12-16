package auth

import (
	"context"
	"medsecurity/type/model"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	PatientRegistration(ctx context.Context, param model.Patient) error
	DoctorRegistration(ctx context.Context, param model.Doctor) error
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

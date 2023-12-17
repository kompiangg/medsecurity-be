package patient_image

import (
	"context"
	"medsecurity/type/model"
	"medsecurity/type/params"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Find(ctx context.Context, param params.RepositoryFindPatientImage) ([]model.PatientImage, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

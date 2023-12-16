package patient

import (
	"context"
	"medsecurity/type/model"
	"medsecurity/type/params"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindPatientByEmail(ctx context.Context, param params.RepoFindPatientByEmailParam) (model.Patient, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

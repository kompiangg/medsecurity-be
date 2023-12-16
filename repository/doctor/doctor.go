package doctor

import (
	"context"
	"medsecurity/type/model"
	"medsecurity/type/params"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindDoctorByEmail(ctx context.Context, param params.RepoFindDoctorByEmailParam) (model.Doctor, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

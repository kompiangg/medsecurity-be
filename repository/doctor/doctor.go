package doctor

import (
	"context"
	pkgSqlx "medsecurity/pkg/db/sqlx"
	"medsecurity/type/model"
	"medsecurity/type/params"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Insert(ctx context.Context, param model.Doctor) (pkgSqlx.Tx, error)
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

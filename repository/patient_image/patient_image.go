package patient_image

import (
	"context"
	pkgSqlx "medsecurity/pkg/db/sqlx"
	"medsecurity/type/model"
	"medsecurity/type/params"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Find(ctx context.Context, param params.RepositoryFindPatientImage) ([]model.PatientImage, error)
	Insert(ctx context.Context, param model.PatientImage) (pkgSqlx.Tx, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

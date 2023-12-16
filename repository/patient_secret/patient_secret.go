package patient_secret

import (
	"context"
	"medsecurity/type/model"

	pkgSqlx "medsecurity/pkg/db/sqlx"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Insert(ctx context.Context, patientSecret model.PatientSecret) (pkgSqlx.Tx, error)
	FindByPatientID(ctx context.Context) (model.PatientSecret, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

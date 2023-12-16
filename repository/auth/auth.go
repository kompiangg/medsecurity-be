package auth

import (
	"context"
	"medsecurity/type/model"

	pkgSqlx "medsecurity/pkg/db/sqlx"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	PatientRegistration(ctx context.Context, param model.Patient) (pkgSqlx.Tx, error)
	DoctorRegistration(ctx context.Context, param model.Doctor) (pkgSqlx.Tx, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

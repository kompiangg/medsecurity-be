package patient

import (
	"context"
	"medsecurity/type/model"
	"medsecurity/type/params"

	pkgSqlx "medsecurity/pkg/db/sqlx"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Insert(ctx context.Context, param model.Patient) (pkgSqlx.Tx, error)
	FindPatientByEmail(ctx context.Context, param params.RepoFindPatientByEmailParam) (model.Patient, error)
	DeleteByID(ctx context.Context, id uuid.UUID) (pkgSqlx.Tx, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

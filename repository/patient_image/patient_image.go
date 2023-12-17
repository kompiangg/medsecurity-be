package patient_image

import (
	"context"
	pkgSqlx "medsecurity/pkg/db/sqlx"
	"medsecurity/type/model"
	"medsecurity/type/params"
	"medsecurity/type/result"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	FindAll(ctx context.Context, param params.RepositoryFindAllPatientImage) ([]model.PatientImage, error)
	Insert(ctx context.Context, param model.PatientImage) (pkgSqlx.Tx, error)
	Find(ctx context.Context, param params.RepositoryFindPatientImage) (model.PatientImage, error)
	InsertPatientRequestGetImageToken(ctx context.Context, param params.RepositoryInsertRequestPatientImageToken) error
	FindPatientRequestGetImageToken(ctx context.Context, param params.RepositoryFindRequestPatientImageToken) (result.RepositoryGetRequestPatientImageToken, error)
}

type repository struct {
	db    *sqlx.DB
	redis *redis.Client
}

func New(db *sqlx.DB, redis *redis.Client) Repository {
	return &repository{
		db:    db,
		redis: redis,
	}
}

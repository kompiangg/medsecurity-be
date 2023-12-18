package access_request

import (
	"context"
	"medsecurity/type/model"
	"medsecurity/type/params"
	"medsecurity/type/result"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	InsertNewRequest(ctx context.Context, param model.AccessRequest) error
	InsertRequestToRedis(ctx context.Context, param params.RepositoryInsertRequestToRedis) error
	FindByImageID(ctx context.Context, imageID uuid.UUID) (model.AccessRequest, error)
	GetDoctorPermissionRedis(ctx context.Context, param params.RepositoryGetDoctorPermissionRedis) (result.RepositoryGetDoctorPermission, error)
	FindByAccessID(ctx context.Context, accessID uuid.UUID) (model.AccessRequest, error)
}

type repository struct {
	db    *sqlx.DB
	redis *redis.Client
}

type Config struct {
}

func New(
	config Config,
	db *sqlx.DB,
	redis *redis.Client,
) Repository {
	return &repository{
		db:    db,
		redis: redis,
	}
}

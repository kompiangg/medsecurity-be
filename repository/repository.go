package repository

import (
	"medsecurity/config"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
}

func New(
	config config.Config,
	db *sqlx.DB,
	// redis *redis.Client,
	// cld objstorage.ObjectStorageItf,
) (Repository, error) {
	return Repository{}, nil
}

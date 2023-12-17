package access_history

import (
	"context"
	"medsecurity/type/model"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Insert(ctx context.Context, param model.AccessHistory) error
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

package db

import (
	"context"
	"database/sql"
	"fmt"

	"medsecurity/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type TrxFn func(*sqlx.Tx) error

func TrxWrapper(ctx context.Context, db *sqlx.DB, name string, opts *sql.TxOptions, fn TrxFn) (err error) {
	tx, err := db.BeginTxx(ctx, opts)
	if err != nil {
		return errors.ErrBeginTransaction
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			err = errors.New("panic happened because: " + fmt.Sprintf("%v", p))
		} else if err != nil {
			if errors.Is(ctx.Err(), context.Canceled) {
				err = errors.Wrap(err, fmt.Sprintf("[%s] Failed to execute tx, cause context was cancelled", name))
			}
		}
	}()

	if err = fn(tx); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, fmt.Sprintf("[%s] Failed to rollback tx", name))
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return errors.ErrCommitTransaction
	}

	return nil
}

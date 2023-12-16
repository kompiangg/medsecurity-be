package patient

import (
	"context"
	"medsecurity/pkg/db/sqlx"
	"medsecurity/pkg/errors"

	"github.com/google/uuid"
)

func (r repository) DeleteByID(ctx context.Context, id uuid.UUID) (sqlx.Tx, error) {
	statement := `
		DELETE FROM patients
		WHERE id = ?
	`

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return tx, errors.Wrap(err, "error at DeleteByID")
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(statement), id)
	if err != nil {
		tx.Rollback()
		return tx, errors.Wrap(err, "error at delete patient")
	}

	return tx, nil
}

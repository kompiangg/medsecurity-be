package auth

import (
	"context"
	"medsecurity/pkg/db/sqlx"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
)

func (r repository) DoctorRegistration(ctx context.Context, param model.Doctor) (sqlx.Tx, error) {
	statement := `
		INSERT INTO doctors (
			id,
			email,
			password,
			full_name
		) VALUES (
			?, ?, 
			?, ?
		)
	`

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return tx, errors.Wrap(err, "error at begin transaction")
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(statement), param.ID, param.Email, param.Password, param.FullName)
	if err != nil {
		return tx, errors.Wrap(err, "error at exec sql")
	}

	return tx, nil
}

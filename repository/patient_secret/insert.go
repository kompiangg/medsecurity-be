package patient_secret

import (
	"context"
	"medsecurity/pkg/db/sqlx"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
)

func (r repository) Insert(ctx context.Context, param model.PatientSecret) (sqlx.Tx, error) {
	statement := `
		INSERT INTO patient_secrets (
			id,
			patient_id,
			private_key,
			public_key,
			key_size,
			salt,
			is_valid
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return tx, errors.Wrap(err, "error at Insert")
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(statement),
		param.ID, param.PatientID,
		param.PrivateKey, param.PublicKey,
		param.KeySize, param.Salt, param.IsValid,
	)

	if err != nil {
		tx.Rollback()
		return tx, errors.Wrap(err, "error at insert patient secret")
	}

	return tx, nil
}

package patient_secret

import (
	"context"
	"database/sql"
	"medsecurity/pkg/db/sqlx"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
)

func (r repository) Insert(ctx context.Context, patientSecret model.PatientSecret) (sqlx.Tx, error) {
	statement := `
		INSERT INTO patient_secrets (
			id,
			patient_id,
			private_key,
			key_size,
			is_valid
		) VALUES (?, ?, ?, ?, ?)
	`

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return tx, errors.Wrap(err, "error at Insert")
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(statement),
		patientSecret.ID,
		patientSecret.PatientID, patientSecret.PrivateKey,
		patientSecret.KeySize, patientSecret.IsValid,
	)

	if err != nil {
		tx.Rollback()
		return tx, errors.Wrap(err, "error at insert patient secret")
	}

	return tx, nil
}

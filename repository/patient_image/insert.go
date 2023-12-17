package patient_image

import (
	"context"
	"medsecurity/pkg/db/sqlx"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
)

func (r repository) Insert(ctx context.Context, param model.PatientImage) (sqlx.Tx, error) {
	q := `
		INSERT INTO patient_images
			(id, patient_id, doctor_id, name, type, url, is_valid)
		VALUES
			(?, ?, ?, ?, ?, ?, ?);
	`

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to begin transaction")
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(q),
		param.ID, param.PatientID, param.DoctorID,
		param.Name, param.Type, param.URL, param.IsValid,
	)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "failed to insert patient image")
	}

	return tx, nil
}

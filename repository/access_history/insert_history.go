package access_history

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
)

func (r repository) Insert(ctx context.Context, param model.AccessHistory) error {
	query := `
		INSERT INTO access_histories
			(id, patient_image_id, patient_id, doctor_id, purpose)
		VALUES
			(?, ?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(
		ctx, r.db.Rebind(query),
		param.ID, param.PatientID, param.DoctorID, param.Purpose,
	)
	if err != nil {
		return errors.Wrap(err, "failed to insert access history")
	}

	return nil
}

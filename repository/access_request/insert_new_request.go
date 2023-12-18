package access_request

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
)

func (r repository) InsertNewRequest(ctx context.Context, param model.AccessRequest) error {
	q := `
		INSERT INTO access_requests (
			id,
			patient_id,
			doctor_id,
			image_id,
			purpose,
			is_allowed,
			allowed_until
		) VALUES (
			?, ?, ?, ?, ?, ?, ?
		)
	`

	_, err := r.db.ExecContext(ctx, r.db.Rebind(q),
		param.ID, param.PatientID, param.DoctorID,
		param.ImageID, param.Purpose, param.IsAllowed, param.AllowedUntil,
	)
	if err != nil {
		return errors.Wrap(err, "error at insert new access request")
	}

	return nil
}

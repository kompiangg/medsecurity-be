package patient_secret

import (
	"context"
	"database/sql"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
)

func (r repository) FindByPatientID(ctx context.Context) (model.PatientSecret, error) {
	var patientSecret model.PatientSecret
	statement := `
		SELECT
			id,
			patient_id,
			private_key,
			key_size,
			is_valid,
			created_at,
			updated_at
		FROM patient_secrets
		WHERE patient_id = ?
	`

	err := r.db.GetContext(ctx, &patientSecret, r.db.Rebind(statement), patientSecret.PatientID)
	if errors.Is(err, sql.ErrNoRows) {
		return model.PatientSecret{}, err
	} else if err != nil {
		return model.PatientSecret{}, errors.Wrap(err, "error at FindByPatientID")
	}

	return patientSecret, nil
}

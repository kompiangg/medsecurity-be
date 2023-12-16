package patient

import (
	"context"
	"database/sql"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
	"medsecurity/type/params"
)

func (r repository) FindPatientByEmail(ctx context.Context, param params.RepoFindPatientByEmailParam) (model.Patient, error) {
	statement := `
		SELECT
			id,
			date_of_birth,
			created_at,
			updated_at,
			password,
			full_name,
			blood_type,
			email,
			phone,
			occupation,
			religion,
			relationship_status,
			nationality,
			address,
			gender
		FROM patients
		WHERE email = ?
	`

	var patient model.Patient
	err := r.db.GetContext(ctx, &patient, r.db.Rebind(statement), param.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return patient, errors.ErrRecordNotFound
	} else if err != nil {
		return patient, errors.Wrap(err, "error at FindPatientByUsername")
	}

	return patient, nil
}

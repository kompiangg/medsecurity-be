package doctor

import (
	"context"
	"database/sql"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
	"medsecurity/type/params"
)

func (r repository) FindDoctorByEmail(ctx context.Context, param params.RepoFindDoctorByEmailParam) (model.Doctor, error) {
	statement := `
		SELECT
			id,
			polyclinic_id,
			email,
			password,
			full_name,
			created_at,
			updated_at
		FROM doctors
		WHERE email = ?
	`

	var doctor model.Doctor
	err := r.db.GetContext(ctx, &doctor, r.db.Rebind(statement), param.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return doctor, err
	} else if err != nil {
		return doctor, errors.Wrap(err, "error at get doctor")
	}

	return doctor, nil
}

package auth

import (
	"context"
	"medsecurity/pkg/db/sqlx"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
)

func (r repository) PatientRegistration(ctx context.Context, param model.Patient) (sqlx.Tx, error) {
	statement := `
		INSERT INTO patients (
			id,
			date_of_birth,
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
		) VALUES (
			?, ?, ?, ?,
			?, ?, ?, ?,
			?, ?, ?, ?,
			?
		)
	`

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return tx, errors.Wrap(err, "error at begin transaction")
	}

	_, err = tx.ExecContext(
		ctx, r.db.Rebind(statement),
		param.ID, param.DateOfBirth, param.Password, param.FullName,
		param.BloodType, param.Email, param.Phone, param.Occupation,
		param.Religion, param.RelationshipStatus, param.Nationality,
		param.Address, param.Gender,
	)
	if err != nil {
		tx.Rollback()
		return tx, errors.Wrap(err, "error at exec sql")
	}

	return tx, nil
}

package patient

import (
	"context"
	"database/sql"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
	"medsecurity/type/params"

	"github.com/Masterminds/squirrel"
)

func (r repository) FindAll(ctx context.Context, param params.RepoFindAllPatients) ([]model.Patient, error) {
	q := squirrel.
		Select(
			"id",
			"date_of_birth",
			"created_at",
			"updated_at",
			"password",
			"full_name",
			"blood_type",
			"email",
			"phone",
			"occupation",
			"religion",
			"relationship_status",
			"nationality",
			"address",
			"gender",
		).From("patients")

	statement, args, err := q.ToSql()
	if err != nil {
		return []model.Patient{}, errors.Wrap(err, "error when build query")
	}

	patient := make([]model.Patient, 0)
	err = r.db.SelectContext(ctx, &patient, r.db.Rebind(statement), args...)
	if errors.Is(err, sql.ErrNoRows) {
		return patient, errors.ErrRecordNotFound
	} else if err != nil {
		return patient, errors.Wrap(err, "error at FindPatientByUsername")
	}

	return patient, nil
}

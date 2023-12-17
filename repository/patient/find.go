package patient

import (
	"context"
	"database/sql"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
	"medsecurity/type/params"

	"github.com/Masterminds/squirrel"
)

func (r repository) Find(ctx context.Context, param params.RepoFindPatient) (model.Patient, error) {
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

	if param.Email.Valid {
		q = q.Where(squirrel.Eq{"email": param.Email})
	}

	if param.ID.Valid {
		q = q.Where(squirrel.Eq{"id": param.ID})
	}

	statement, args, err := q.ToSql()
	if err != nil {
		return model.Patient{}, errors.Wrap(err, "error when build query")
	}

	var patient model.Patient
	err = r.db.GetContext(ctx, &patient, r.db.Rebind(statement), args...)
	if errors.Is(err, sql.ErrNoRows) {
		return patient, errors.ErrRecordNotFound
	} else if err != nil {
		return patient, errors.Wrap(err, "error at FindPatientByUsername")
	}

	return patient, nil
}

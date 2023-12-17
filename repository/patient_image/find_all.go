package patient_image

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
	"medsecurity/type/params"

	"github.com/Masterminds/squirrel"
)

func (r repository) Find(ctx context.Context, param params.RepositoryFindPatientImage) ([]model.PatientImage, error) {
	q := squirrel.Select(
		"id",
		"patient_id",
		"doctor_id",
		"name",
		"type",
		"url",
		"is_valid",
		"created_at",
		"updated_at",
	).
		From("patient_images").
		Limit(param.Limit).
		Offset(param.Offset)

	if param.PatientID.Valid {
		q = q.Where(squirrel.Eq{"patient_id": param.PatientID})
	}

	if param.DoctorID.Valid {
		q = q.Where(squirrel.Eq{"doctor_id": param.DoctorID})
	}

	statement, args, err := q.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "error when build query")
	}

	var patientImages []model.PatientImage
	err = r.db.SelectContext(ctx, &patientImages, r.db.Rebind(statement), args...)
	if err != nil {
		return patientImages, err
	}

	return patientImages, nil
}

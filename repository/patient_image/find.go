package patient_image

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
	"medsecurity/type/params"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s repository) Find(ctx context.Context, param params.RepositoryFindPatientImage) (model.PatientImage, error) {
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
	).From("patient_images").
		Where(squirrel.Eq{"is_valid": param.IsValid})

	if param.ID != uuid.Nil {
		q = q.Where(squirrel.Eq{"id": param.ID})
	}

	if param.PatientID.Valid {
		q = q.Where(squirrel.Eq{"patient_id": param.PatientID})
	}

	if param.DoctorID.Valid {
		q = q.Where(squirrel.Eq{"doctor_id": param.DoctorID})
	}

	statement, args, err := q.ToSql()
	if err != nil {
		return model.PatientImage{}, errors.Wrap(err, "error when build query")
	}

	var patientImage model.PatientImage
	err = s.db.GetContext(ctx, &patientImage, s.db.Rebind(statement), args...)
	if errors.Is(err, errors.ErrRecordNotFound) {
		return patientImage, errors.ErrRecordNotFound
	} else if err != nil {
		return patientImage, errors.Wrap(err, "error when get patient image")
	}

	return patientImage, nil
}

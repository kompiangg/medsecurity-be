package doctor

import (
	"context"
	"database/sql"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
	"medsecurity/type/params"

	"github.com/Masterminds/squirrel"
)

func (r repository) Find(ctx context.Context, param params.RepoFindDoctor) (model.Doctor, error) {
	q := squirrel.Select(
		"id",
		"polyclinic_id",
		"email",
		"password",
		"full_name",
		"created_at",
		"updated_at",
	).From("doctors")

	if param.Email.Valid {
		q = q.Where(squirrel.Eq{"email": param.Email})
	}

	if param.ID.Valid {
		q = q.Where(squirrel.Eq{"id": param.ID})
	}

	statement, args, err := q.ToSql()
	if err != nil {
		return model.Doctor{}, errors.Wrap(err, "error when build query")
	}

	var doctor model.Doctor
	err = r.db.GetContext(ctx, &doctor, r.db.Rebind(statement), args...)
	if errors.Is(err, sql.ErrNoRows) {
		return doctor, errors.ErrRecordNotFound
	} else if err != nil {
		return doctor, errors.Wrap(err, "error at get doctor")
	}

	return doctor, nil
}

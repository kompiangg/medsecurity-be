package access_request

import (
	"context"
	"database/sql"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"

	"github.com/google/uuid"
)

func (r repository) FindByImageID(ctx context.Context, imageID uuid.UUID) (model.AccessRequest, error) {
	q := `
		SELECT
			id,
			patient_id,
			doctor_id,
			image_id,
			purpose,
			is_allowed,
			allowed_until
		FROM access_requests
		WHERE
			image_id = ?
		ORDER BY
			created_at DESC
	`

	var res model.AccessRequest
	err := r.db.GetContext(ctx, &res, r.db.Rebind(q), imageID)
	if errors.Is(err, sql.ErrNoRows) {
		return res, errors.ErrRecordNotFound
	} else if err != nil {
		return model.AccessRequest{}, errors.Wrap(err, "error at find access request by image id")
	}

	return res, nil
}

package access_request

import (
	"context"
	"medsecurity/pkg/errors"

	"github.com/google/uuid"
)

func (r repository) DisallowRequestByID(ctx context.Context, id uuid.UUID) error {
	q := `
		UPDATE access_requests
		SET
			is_allowed=false
		WHERE
			id = ?
	`

	_, err := r.db.ExecContext(ctx, r.db.Rebind(q), id)
	if err != nil {
		return errors.Wrap(err, "error at disallow request by id")
	}

	return nil
}

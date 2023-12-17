package cloudinary

import (
	"context"
	"medsecurity/pkg/errors"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (r repository) Remove(ctx context.Context, pubclicID string) error {
	_, err := r.cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: pubclicID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to remove file")
	}

	return nil
}

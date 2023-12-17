package cloudinary

import (
	"context"
	"io"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (c repository) UploadEncryptedFile(ctx context.Context, content io.Reader) (uploader.UploadResult, error) {
	resp, err := c.cloudinary.Upload.Upload(ctx, content, uploader.UploadParams{
		UniqueFilename: api.Bool(true),
		Overwrite:      api.Bool(false),
		ResourceType:   "auto",
	})
	if err != nil {
		return uploader.UploadResult{}, err
	}

	return *resp, nil
}

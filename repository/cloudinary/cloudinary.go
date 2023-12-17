package cloudinary

import (
	"context"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Repository interface {
	UploadEncryptedFile(ctx context.Context, content io.Reader) (uploader.UploadResult, error)
	Remove(ctx context.Context, publicID string) error
}

type repository struct {
	cloudinary *cloudinary.Cloudinary
}

type Config struct {
	URIConnection string
}

func New(config Config) (Repository, error) {
	cld, err := cloudinary.NewFromURL(config.URIConnection)
	if err != nil {
		return nil, err
	}

	return &repository{
		cloudinary: cld,
	}, nil
}

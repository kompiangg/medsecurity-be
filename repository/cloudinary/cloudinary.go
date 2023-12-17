package cloudinary

import (
	"context"
	"io"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Repository interface {
	UploadEncryptedFile(ctx context.Context, content io.Reader) (uploader.UploadResult, error)
	Remove(ctx context.Context, publicID string) error
	DownloadFile(ctx context.Context, url string) ([]byte, error)
}

type repository struct {
	cloudinary *cloudinary.Cloudinary
	httpClient *http.Client
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
		httpClient: &http.Client{},
	}, nil
}

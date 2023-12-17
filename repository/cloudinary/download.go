package cloudinary

import (
	"context"
	"io"
	"medsecurity/pkg/errors"
)

func (c repository) DownloadFile(ctx context.Context, url string) ([]byte, error) {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error when downloading file")
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, errors.ErrNotFound
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error when reading response body")
	}

	return resBody, nil
}

package patient_image

import (
	"context"
	"fmt"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/type/result"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r repository) InsertPatientRequestGetImageToken(ctx context.Context, param params.RepositoryInsertRequestPatientImageToken) error {
	err := r.redis.HSet(ctx, fmt.Sprintf(PatientRequestGetImage, param.PatientID), param).Err()
	if err != nil {
		return errors.Wrap(err, "error when set token")
	}

	err = r.redis.Expire(ctx, fmt.Sprintf(PatientRequestGetImage, param.PatientID), time.Minute*time.Duration(param.ValidInMinute)).Err()
	if err != nil {
		return errors.Wrap(err, "error when set expire")
	}

	return nil
}

func (r repository) FindPatientRequestGetImageToken(ctx context.Context, param params.RepositoryFindRequestPatientImageToken) (result.RepositoryGetRequestPatientImageToken, error) {
	res := result.RepositoryGetRequestPatientImageToken{}

	err := r.redis.HGetAll(ctx, fmt.Sprintf(PatientRequestGetImage, param.PatientID)).Scan(&res)
	if errors.Is(err, redis.Nil) {
		return res, errors.ErrNotFound
	} else if err != nil {
		return res, errors.Wrap(err, "error when get token")
	}

	// err = r.redis.Del(ctx, fmt.Sprintf(PatientRequestGetImage, param.PatientID)).Err()
	// if err != nil {
	// 	return res, errors.Wrap(err, "error when delete token")
	// }

	return res, nil
}

package access_request

import (
	"context"
	"fmt"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/type/result"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r repository) InsertRequestToRedis(ctx context.Context, param params.RepositoryInsertRequestToRedis) error {
	key := fmt.Sprintf("access_request:%s:%s", param.DoctorID, param.RequestID)

	err := r.redis.Del(ctx, fmt.Sprintf("access_request:%s:*", param.DoctorID)).Err()
	if err != nil {
		return errors.Wrap(err, "error at delete access request to redis")
	}

	err = r.redis.HSet(ctx, key, param).Err()
	if err != nil {
		return errors.Wrap(err, "error at set access request to redis")
	}

	err = r.redis.Expire(ctx, key, time.Hour*24*time.Duration(param.KeepAliveInDays)).Err()
	if err != nil {
		return errors.Wrap(err, "error at set access request to redis")
	}

	return nil
}

func (r repository) GetDoctorPermissionRedis(ctx context.Context, param params.RepositoryGetDoctorPermissionRedis) (result.RepositoryGetDoctorPermission, error) {
	key := fmt.Sprintf("access_request:%s:%s", param.DoctorID, param.RequestID)

	res := result.RepositoryGetDoctorPermission{}
	err := r.redis.HGetAll(ctx, key).Scan(&res)
	if errors.Is(err, redis.Nil) {
		return result.RepositoryGetDoctorPermission{}, errors.ErrRecordNotFound
	} else if err != nil {
		return result.RepositoryGetDoctorPermission{}, errors.Wrap(err, "error at get doctor permission redis")
	}

	if res.Password == "" {
		return result.RepositoryGetDoctorPermission{}, errors.ErrRecordNotFound
	}

	return res, nil
}

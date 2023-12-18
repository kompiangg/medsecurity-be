package access_request

import (
	"context"
	"fmt"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"time"
)

func (r repository) InsertRequestToRedis(ctx context.Context, param params.RepositoryInsertRequestToRedis) error {
	key := fmt.Sprintf("access_request:%s", param.RequestID)

	err := r.redis.HSet(ctx, key, param).Err()
	if err != nil {
		return errors.Wrap(err, "error at set access request to redis")
	}

	err = r.redis.Expire(ctx, key, time.Hour*24*time.Duration(param.KeepAliveInDays)).Err()
	if err != nil {
		return errors.Wrap(err, "error at set access request to redis")
	}

	return nil
}

package ratelimit

import (
	"context"
	"github.com/go-faster/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type (
	redisRateLimitRepository struct {
		redisClient *redis.Client
		limit       int64         // 제한 횟수
		expiration  time.Duration // 만기 기한
	}

	RedisRateLimitRepository interface {
		GetRateLimitCount(
			ctx context.Context,
			key string,
		) (int64, error)
		IncrementRateLimitCount(
			ctx context.Context,
			key string,
		) error
	}
)

func NewRedisRateLimitRepository(
	redisClient *redis.Client,
	limit int64,
	expiration time.Duration,
) RedisRateLimitRepository {
	return &redisRateLimitRepository{
		redisClient: redisClient,
		limit:       limit,
		expiration:  expiration,
	}
}

func (r *redisRateLimitRepository) GetRateLimitCount(
	ctx context.Context,
	key string,
) (int64, error) {
	rateLimitCount, err := r.redisClient.Get(ctx, key).Int64()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil // 키가 존재하지 않으면 기본값 0 반환
		}

		return 0, errors.Wrapf(err, "레디스 키 조회 중 오류가 발생했습니다")
	}

	return rateLimitCount, nil
}

// IncrementRateLimitCount
// (ex) 1분에 3회 제한 => limit : 3, expiration : Time.minute*1
// (ex) 하루에 10회 제한 => limit : 10, expiration : Time.Day*1
func (r *redisRateLimitRepository) IncrementRateLimitCount(
	ctx context.Context,
	key string,
) error {
	rateLimitCount, err := r.GetRateLimitCount(ctx, key)

	if err != nil {
		return err
	}

	if rateLimitCount >= r.limit { // 현재 카운트가 제한을 초과한 경우 더 이상 카운트를 올리지 않습니다.
		return nil
	}

	rateLimitCount, err = r.redisClient.Incr(ctx, key).Result()

	if err != nil {
		return errors.Wrapf(err, "레디스 키 증가 중 오류가 발생했습니다")
	}

	if rateLimitCount == 1 { // 처음 increment한 경우, 만료 기한을 지정합니다.

		err = r.redisClient.Expire(ctx, key, r.expiration).Err()

		if err != nil {
			return errors.Wrapf(err, "레디스 키 만료 설정 중 오류가 발생했습니다")
		}
	}

	return nil
}

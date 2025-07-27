package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
)

type RedisCacheService struct {
	ctx context.Context
	rdb *redis.Client
}

func NewRedisCacheService(rdb *redis.Client) adapters.ICacheRepository {
	return &RedisCacheService{
		ctx: context.Background(),
		rdb: rdb,
	}
}

func (cs *RedisCacheService) Set(key string, value interface{}, ttl time.Duration) error {
	err := cs.rdb.Set(cs.ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return cs.rdb.Set(cs.ctx, key, value, ttl).Err()
}

func (cs *RedisCacheService) Get(key string, dest any) error {
	data, err := cs.rdb.Get(cs.ctx, key).Result()

	if err == redis.Nil {
		return err
	}

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

func (cs *RedisCacheService) Clear(pattern string) error {
	cursor := uint64(0)

	for {
		keys, nextCursor, err := cs.rdb.Scan(cs.ctx, cursor, pattern, 2).Result()
		if err != nil {
			return err
		}

		if len(keys) > 0 {
			cs.rdb.Del(cs.ctx, keys...)
		}

		cursor = nextCursor

		if cursor == 0 {
			break
		}
	}

	return nil
}

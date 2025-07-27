package adapters

import "time"

type ICacheRepository interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string, dest any) error
	Clear(pattern string) error
}

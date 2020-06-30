package base

import "time"

type (
	Cache interface {
		SetPrefix(prefix string)
		Get(key string) (data interface{}, err error)
		Set(key string, val interface{}, expire time.Duration) bool
		Rm(key string) bool
		Has(key string) bool
		Expire(key string, expire time.Duration) bool
	}
)

const (
	DefaultCachePrefix = ""
	AccessTokenKey     = "buffgeWechatAccessToken"
)

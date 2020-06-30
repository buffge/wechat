package cache

import (
	"errors"
	"github.com/buffge/wechat/pkg/base"
	goCache "github.com/patrickmn/go-cache"
	"time"
)

var defaultCache *DefaultCache

type (
	DefaultCache struct {
		c      *goCache.Cache
		prefix string
	}
)

func newDefaultCache() *DefaultCache {
	return &DefaultCache{
		c:      goCache.New(0, 2*time.Minute),
		prefix: base.DefaultCachePrefix,
	}
}
func GetDefaultCache() *DefaultCache {
	if defaultCache == nil {
		defaultCache = newDefaultCache()
	}
	return defaultCache
}

func (d *DefaultCache) SetPrefix(prefix string) {
	d.prefix = prefix
}
func (d *DefaultCache) Get(key string) (data interface{}, err error) {
	key = d.prefix + key
	var exist bool
	if data, exist = d.c.Get(key); !exist {
		return nil, errors.New("not found")
	}
	return data, nil
}

func (d *DefaultCache) Set(key string, val interface{}, expire time.Duration) bool {
	key = d.prefix + key
	d.c.Set(key, val, expire)
	return true
}

func (d *DefaultCache) Rm(key string) bool {
	key = d.prefix + key
	d.c.Delete(key)
	return true
}

func (d *DefaultCache) Has(key string) bool {
	key = d.prefix + key
	_, exist := d.c.Get(key)
	return exist
}

func (d *DefaultCache) Expire(key string, expire time.Duration) bool {
	key = d.prefix + key
	if v, err := d.Get(key); err != nil {
		return false
	} else {
		d.Set(key, v, expire)
	}
	return true
}

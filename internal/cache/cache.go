
package cache

import (
    "gocool/internal/model"
    "time"
    "github.com/patrickmn/go-cache"
)

// CacheService provides caching functionality for the application
type CacheService struct {
    Cache *cache.Cache
}

// NewCacheService creates a new instance of CacheService with default settings
func NewCacheService(defaultExpiration, cleanupInterval time.Duration) *CacheService {
    return &CacheService{
        Cache: cache.New(defaultExpiration, cleanupInterval),
    }
}

// GetUserFromCache retrieves a user from the cache
func (cs *CacheService) GetUserFromCache(key string) (model.User, bool) {
    if x, found := cs.Cache.Get(key); found {
        return x.(model.User), true
    }
    return model.User{}, false
}

// SetUserInCache adds a user to the cache
func (cs *CacheService) SetUserInCache(key string, user model.User) {
    cs.Cache.Set(key, user, cache.DefaultExpiration)
}

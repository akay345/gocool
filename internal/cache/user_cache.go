package cache

import (
    "gocool/internal/model"
)

// GetUserDetails fetches the user details from the cache, or loads it from the database if not present
// shouldn't this be in reverse order.
// getUserDetails should be implemented in repository and inside repository first fetch from cache and if not
// found get from db and refill the cache.
// this way we are maintaining things in 2 places, repository and cache. this cache implementation can be avoided.
func (cs *CacheService) GetUserDetails(userID int64) (model.User, bool) {
    cacheKey := "user:" + string(userID)
    user, found := cs.GetUserFromCache(cacheKey)
    if !found {
        // Load from database
        user, err := database.GetUserByID(userID)
        if err != nil {
            return model.User{}, false
        }
        cs.SetUserInCache(cacheKey, user)
        return user, true
    }
    return user, true
}

// InvalidateUserCache invalidates the cache for a user
func (cs *CacheService) InvalidateUserCache(userID int64) {
    cacheKey := "user:" + string(userID)
    cs.Cache.Delete(cacheKey)
}

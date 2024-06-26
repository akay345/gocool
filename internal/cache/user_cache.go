
package cache

import (
    "gocool/internal/model"
)

// GetUserDetails fetches the user details from the cache, or loads it from the database if not present
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

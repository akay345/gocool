
package service

import (
    "gocool/internal/model"
)

// IUserService defines the interface for user service operations
type IUserService interface {
    RegisterUser(user model.User) error
    UpdateUserDetails(user model.User) error
    DeleteUser(id int64) error
}

package users

import (
	"gocool/internal/model"
)

// IUserService defines the interface for user service operations
type IEnachService interface {
	EnachCreate(user model.User) error
	EnachDelete(user model.User) error
	EnachGetDetails(id int64) error
}

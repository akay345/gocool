package services

import (
	"errors"
	"gocool/internal/model"
	"gocool/internal/repository"
)

// UserService handles the business logic for user management
type EnachService struct {
	Repo *repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo *repository.UserRepository) *EnachService {
	return &EnachService{Repo: repo}
}

// RegisterUser handles the registration of a new user
func (e *EnachService) EnachCreate(user model.User) error {
	// Check if user already exists
	_, err := e.Repo.FindUserByEmail(user.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Create a new user
	err = model.CreateUser(e.Repo.DB, user)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserDetails updates the information of an existing user
func (s *EnachService) EnachDelete(user model.User) error {
	return model.UpdateUser(s.Repo.DB, user)
}

// DeleteUser handles the deletion of a user
func (s *EnachService) EnachGetDetails(id int64) error {
	return model.DeleteUser(s.Repo.DB, id)
}

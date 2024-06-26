
package service

import (
    "gocool/internal/model"
    "gocool/internal/repository"
    "errors"
)

// UserService handles the business logic for user management
type UserService struct {
    Repo *repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

// RegisterUser handles the registration of a new user
func (s *UserService) RegisterUser(user model.User) error {
    // Check if user already exists
    _, err := s.Repo.FindUserByEmail(user.Email)
    if err == nil {
        return errors.New("user already exists")
    }

    // Create a new user
    err = model.CreateUser(s.Repo.DB, user)
    if err != nil {
        return err
    }
    return nil
}

// UpdateUserDetails updates the information of an existing user
func (s *UserService) UpdateUserDetails(user model.User) error {
    return model.UpdateUser(s.Repo.DB, user)
}

// DeleteUser handles the deletion of a user
func (s *UserService) DeleteUser(id int64) error {
    return model.DeleteUser(s.Repo.DB, id)
}

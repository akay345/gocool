package controller

import (
	"encoding/json"
	"gocool/internal/model"
	"gocool/internal/service/users"
	"net/http"
	"strconv"
)

// UserController handles web requests for user operations
type UserController struct {
	UserService *users.UserService
}

// NewUserController creates a new instance of UserController
func NewUserController(userService *users.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// RegisterUser handles the user registration process
func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := uc.UserService.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User registered successfully")
}

// GetUserDetails handles fetching a user by ID
func (uc *UserController) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := uc.UserService.GetUserDetails(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

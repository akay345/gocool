
package repository

import (
    "database/sql"
    "gocool/internal/model"
)

// UserRepository handles the more complex database operations for the user model
type UserRepository struct {
    DB *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{DB: db}
}

// FindUserByEmail fetches a user by their email address
func (repo *UserRepository) FindUserByEmail(email string) (model.User, error) {
    var user model.User
    query := \`SELECT id, username, email, password, created_at, updated_at FROM users WHERE email = ?\`
    err := repo.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
    return user, err
}

// GetAllUsers retrieves all users from the database
func (repo *UserRepository) GetAllUsers() ([]model.User, error) {
    query := \`SELECT id, username, email, password, created_at, updated_at FROM users\`
    rows, err := repo.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []model.User
    for rows.Next() {
        var user model.User
        if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

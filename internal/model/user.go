
package model

import (
    "database/sql"
    "time"
)

// User represents the user model in the database
type User struct {
    ID        int64     `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB, user User) error {
    query := \`INSERT INTO users (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)\`
    _, err := db.Exec(query, user.Username, user.Email, user.Password, time.Now(), time.Now())
    return err
}

// GetUserByID fetches a user by ID from the database
func GetUserByID(db *sql.DB, id int64) (User, error) {
    var user User
    query := \`SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?\`
    err := db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
    return user, err
}

// UpdateUser updates an existing user in the database
func UpdateUser(db *sql.DB, user User) error {
    query := \`UPDATE users SET username = ?, email = ?, updated_at = ? WHERE id = ?\`
    _, err := db.Exec(query, user.Username, user.Email, time.Now(), user.ID)
    return err
}

// DeleteUser removes a user from the database
func DeleteUser(db *sql.DB, id int64) error {
    query := \`DELETE FROM users WHERE id = ?\`
    _, err := db.Exec(query, id)
    return err
}

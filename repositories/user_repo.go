package repositories

import (
	"context"
	"database/sql"
	"splitwise-backend/db"
	"splitwise-backend/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	GetByEmail(email string) (*models.User, error) // Updated method signature
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (name, email, password_hash, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.DB.QueryRow(context.Background(), query, user.Name, user.Email, user.PasswordHash, user.CreatedAt).Scan(&user.ID)
	return err
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, password_hash, created_at FROM users WHERE id = $1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET name = $1, email = $2, password_hash = $3 WHERE id = $4`
	_, err := db.DB.Exec(context.Background(), query, user.Name, user.Email, user.PasswordHash, user.ID)
	return err
}

func (r *userRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.DB.Exec(context.Background(), query, id)
	return err
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) { // Updated method implementation
	var user models.User
	query := `SELECT id, name, email, password_hash FROM users WHERE email = $1`
	err := db.DB.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err // Database error
	}
	return &user, nil // User found
}

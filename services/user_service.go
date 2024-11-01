package services

import (
	"fmt"
	"log"
	"splitwise-backend/models"
	"splitwise-backend/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	ValidateUser(email string, password string) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	if user.PasswordHash != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.PasswordHash = string(passwordHash)
	}
	return s.userRepo.UpdateUser(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}

func (s *userService) CreateUser(user *models.User) error {
	log.Println("Raw Password:", user.Password) // Debug: log raw password input

	// Hash the plaintext password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	user.PasswordHash = string(hashedPassword)    // Save hashed password only
	log.Println("Raw Password 2:", user.Password) // Debug: log raw password input
	user.Password = ""                            // Clear plaintext password

	log.Printf("Hashed Password (Signup): %s", user.PasswordHash) // Add logging
	log.Println("Raw Password 3:", user.Password)                 // Debug: log raw password input

	// Save user in the database
	return s.userRepo.CreateUser(user)
}

func (s *userService) ValidateUser(email string, password string) (*models.User, error) {
	// Fetch user by email
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	// Compare provided password with stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	// Clear password hash before returning user
	user.PasswordHash = ""
	return user, nil
}

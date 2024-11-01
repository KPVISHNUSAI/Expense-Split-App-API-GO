package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"splitwise-backend/models"
	"splitwise-backend/services"
	"splitwise-backend/utils"
	"time"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	userService services.UserService
}

func NewAuthHandler(userService services.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user := &models.User{
		Name:      credentials.Username,
		Email:     credentials.Email,
		Password:  credentials.Password, // this should be used for hashing
		CreatedAt: time.Now(),           // Assuming you have a CreatedAt field
	}

	err := h.userService.CreateUser(user)
	if err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	log.Printf("Login attempt for email: %s", credentials.Email)

	user, err := h.userService.ValidateUser(credentials.Email, credentials.Password)

	if err != nil {
		log.Printf("Login error: %v", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(credentials.Email)
	if err != nil {
		log.Printf("Token generation error: %v", err)
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"email": user.Email,
		"name":  user.Name,
	})
}

func RegisterAuthRoutes(router *mux.Router, userService services.UserService) {
	handler := NewAuthHandler(userService)
	router.HandleFunc("/api/signup", handler.Signup).Methods("POST")
	router.HandleFunc("/api/login", handler.Login).Methods("POST")
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/wisnu-bdd/be-go-with-auth-template/config"
	"github.com/wisnu-bdd/be-go-with-auth-template/middleware"
	"github.com/wisnu-bdd/be-go-with-auth-template/models"
	"github.com/wisnu-bdd/be-go-with-auth-template/services"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDetailsRequest struct {
	Email string `json:"email"`
}

type UpdateUserPasswordRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type PublicUser struct {
	Email           string `json:"email"`
	Name string `json:"name"`
	// Add other safe fields here if needed
}

type MultipleUsers struct {
	Data []*models.User `json:"data"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Password != req.ConfirmPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create User struct
	newUser := &models.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert user
	err = services.InsertUser(newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Retrieve user from DB
	user, err := services.GetUserByEmail(req.Email)
	if err != nil {
		log.Println("Login error:", err)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized) // Generic message
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Create JWT claims
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWTSecret)
	if err != nil {
		log.Println("JWT signing error:", err)
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token":   tokenString,
		"email":   user.Email, // take out later
		"message": "Login successful",
	})
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	email := middleware.GetUserEmail(r)
	if email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := services.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	publicUser := PublicUser{
		Email: email,
		Name: user.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publicUser)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, "Failed to fetch menuGroups", http.StatusInternalServerError)
		return
	}

	if users == nil {
		users = []*models.User{}
	}

	response := MultipleUsers{
		Data: users,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateUserDetailsByEmail(w http.ResponseWriter, r *http.Request) {
	email := middleware.GetUserEmail(r)
	if email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UpdateUserDetailsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create User struct
	newUser := &models.User{
		UpdatedAt: time.Now(),
	}

	// Insert user
	err := services.UpdateUserDetailsByEmail(email, newUser)
	if err != nil {
		http.Error(w, "Failed to update user details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User updated successfully",
	})
}

func UpdateUserPasswordByEmail(w http.ResponseWriter, r *http.Request) {
	email := middleware.GetUserEmail(r)
	if email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UpdateUserPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Password != req.ConfirmPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create User struct
	newUser := &models.User{
		Password:  string(hashedPassword),
		UpdatedAt: time.Now(),
	}

	// Insert user
	err = services.UpdateUserPasswordByEmail(email, newUser)
	if err != nil {
		http.Error(w, "Failed to update user password", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User password updated successfully",
	})
}

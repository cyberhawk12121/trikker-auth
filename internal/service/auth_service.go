package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/trikker/auth/internal/models"
	"github.com/trikker/auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo      *repository.SQLUserRepository
	jwtSecret string
}

type UserRegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	userId int
	Token  string `json:"token"`
}

func NewAuthService(repo *repository.SQLUserRepository, jwtSecret string) *AuthService {
	return &AuthService{repo, jwtSecret}
}

func (s *AuthService) RegisterUser(c *gin.Context) (*TokenResponse, error) {
	// Get the body from the request
	// Parse the body into a user struct
	var user UserRegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	// Check if the provided email is already registered
	existingUser, err := s.repo.FindUserByEmail(user.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("email already registered")
	}
	// Run validations on the user input

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create a new user
	newUser := &models.User{
		First_Name: user.FirstName,
		Last_Name:  user.LastName,
		Email:      user.Email,
		Password:   string(hashedPassword),
	}

	// Save the user to the database
	createdUser, err := s.repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	// Create a session token
	tokenString, err := s.generateJWT(newUser)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{int(createdUser.ID), tokenString}, nil
}

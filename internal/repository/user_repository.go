package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/trikker/auth/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id int64) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int64) error
}

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db}
}

func (s SQLUserRepository) CreateUser(user *models.User) (*models.User, error) {
	fmt.Println("Creating user", user)
	query := `
	INSERT INTO users (first_name, last_name, password_hash, email, user_type, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`

	now := time.Now()
	err := s.db.QueryRow(query, user.First_Name, user.Last_Name, user.Password, user.Email, "normal", now, now).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

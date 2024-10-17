package models

import (
	"time"
)

type User struct {
	ID         int64     `json:"id"`
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Password   string    `json:"-"`
	Email      string    `json:"email"`
	User_Type  string    `json:"user_type"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

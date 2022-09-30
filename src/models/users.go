package models

import (
	"time"
)

type User struct {
	ID          string         `json:"id" gorm:"type:uuid;primary_key"`
	Name        string         `json:"name"`
	Email  			string         `json:"email"`
	Password  	string     		 `json:"password"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
}
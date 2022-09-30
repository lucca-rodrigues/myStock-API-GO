package models

import (
	"time"
)

type User struct {
	// ID          uint           `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	ID          string         `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string         `json:"name"`
	Email  			string         `json:"email"`
	Password  	string         `json:"password"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`

}
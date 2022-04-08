package entity

import (
	"time"
)

// Users struct
type Users struct {
	ID             int        `gorm:"column:id;primary_key;auto_increment;not null"`
	Firstname      string     `gorm:"column:first_name;"`
	Lastname       string     `gorm:"column:last_name"`
	Username       string     `gorm:"column:user_name;not null"`
	Password       string     `gorm:"column:password;not null"`
	Role           string     `gorm:"column:role"`
	Token          *string    `gorm:"column:token"`
	TokenExpriedAt *time.Time `gorm:"column:token_expried_at"`
	BaseModel
}

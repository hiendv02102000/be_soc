package entity

import (
	"time"
)

// Users struct
type Novels struct {
	ID             int         `gorm:"column:id;primary_key;auto_increment;not null"`
	Name           string      `gorm:"column:name;"`
	ImageUrl       string      `gorm:"column:last_name"`
	View           int         `gorm:"column:username;not null"`
	Role           roleAllowed `gorm:"column:role" sql:"type:role_name"`
	Token          *string     `gorm:"column:token"`
	TokenExpriedAt *time.Time  `gorm:"column:token_expried_at"`
	BaseModel
}

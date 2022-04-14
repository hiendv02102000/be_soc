package entity

// Comments struct
type Comments struct {
	ID              int    `gorm:"column:id;primary_key;auto_increment;not null"`
	CommentsContent string `gorm:"column:cmt_content"`
	UsersId         int    `gorm:"column:users_id"`
	BaseModel
}

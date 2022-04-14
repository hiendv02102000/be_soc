package entity

type UsersNovels struct {
	ID       int `gorm:"column:id;primary_key;auto_increment;not null"`
	UsersID  int `gorm:"column:users_id"`
	NovelsID int `gorm:"column:novels_id"`

	BaseModel
}

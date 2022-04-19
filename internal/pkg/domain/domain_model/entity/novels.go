package entity

type Novels struct {
	ID       int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Name     string `gorm:"column:name;"`
	ImageUrl *string `gorm:"column:image_url"`
	View     int    `gorm:"column:view"`
	UsersID  int    `gorm:"column:users_id"`
	Chapters []Chapters
	BaseModel
}

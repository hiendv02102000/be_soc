package entity

type Chapters struct {
	ID         int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Title      string `gorm:"column:title;"`
	ContentUrl string `gorm:"column:content_url"`
	NovelsID   int    `gorm:"column:novels_id"`
	CommentsID int    `gorm:"column:comments_id"`

	BaseModel
}

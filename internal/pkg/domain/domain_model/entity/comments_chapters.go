package entity

type CommentsChapters struct {
	ID         int `gorm:"column:id;primary_key;auto_increment;not null"`
	CommentsID int `gorm:"column:comments_id"`
	ChaptersID int `gorm:"column:chapters_id"`

	BaseModel
}

package entity

type Chapters struct {
	ID         int    `gorm:"column:id;primary_key;auto_increment;not null" json:"ID"`
	Title      string `gorm:"column:title;" json:"title"`
	ContentUrl string `gorm:"column:content_url" json:"content_url"`
	NovelsID   int    `gorm:"column:novels_id" json:"novel_id"`

	BaseModel
}

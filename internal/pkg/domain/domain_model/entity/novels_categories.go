package entity

type NovelsCategories struct {
	ID           int `gorm:"column:id;primary_key;auto_increment;not null"`
	NovelsID     int `gorm:"column:novels_id"`
	CategoriesID int `gorm:"column:categories_id"`

	BaseModel
}

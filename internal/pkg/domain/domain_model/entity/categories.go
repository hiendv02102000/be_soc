package entity

type Categories struct {
	ID   int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Name string `gorm:"column:name;"`
	BaseModel
}

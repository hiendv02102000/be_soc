package dto

// CreateNovelRequest struct
type CreateNovelRequest struct {
	Name         string `validate:"required"`
	CategoriesID []int  ``
}

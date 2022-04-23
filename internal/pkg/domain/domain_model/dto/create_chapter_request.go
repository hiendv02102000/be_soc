package dto

type CreateChapterRequest struct {
	NovelID int    `validate:"required"`
	Title   string `validate:"required"`
}

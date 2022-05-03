package dto

// LoginResponse struct
type UpdateChapterRequest struct {
	NovelID    int     `json:"novel_id"`
	ChapterID  int     `json:"chapter_id"`
	Title      string  `json:"title"`
	Contenturl *string `json:"contenturl"`
}

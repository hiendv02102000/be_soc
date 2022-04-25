package dto

// CreateNovelRequest struct
type CreateCommentRequest struct {
	CommentContent string `json:"cmt_content" validate:"required"`
	ChapterID      int    `json:"chapter_id" validate:"required"`
}

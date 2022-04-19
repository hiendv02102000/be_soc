package dto

// LoginResponse struct
type NovelListRequest struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Categories    string `json:"categories"`
	UserID        int    `json:"user_id"`
	Isgetchapters bool   `json:"isgetchapter"`
}

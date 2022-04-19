package dto

// LoginResponse struct
type ListNovelsRequest struct {
	Name          string `json:"name"`
	Categories    string `json:"categories"`
	UserID        int    `json:"id"`
	Isgetchapters bool   `json:"isgetchapter"`
}

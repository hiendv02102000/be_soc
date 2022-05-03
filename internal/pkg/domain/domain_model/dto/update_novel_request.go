package dto

// LoginResponse struct
type UpdateNovelRequest struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Imageurl *string `json:"imageurl"`
}

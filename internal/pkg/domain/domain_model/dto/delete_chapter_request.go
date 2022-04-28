package dto

type DeleteChapterRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Imageurl string `json:"imageurl"`
}

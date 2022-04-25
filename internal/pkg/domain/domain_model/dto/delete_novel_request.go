package dto

type DeleteNovelRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Imageurl string `json:"imageurl"`
}

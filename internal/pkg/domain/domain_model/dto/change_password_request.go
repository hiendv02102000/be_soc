package dto

type ChangePasswordRequest struct {
	Password string `json:"password" validate:"required,password"`
}

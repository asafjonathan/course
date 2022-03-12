package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required", regexp=^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	Password string `json:"password" validate:"required"`
}

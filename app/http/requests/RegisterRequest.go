package requests

type RegisterRequest struct {
	Name     string `json:"name" validate:"required" form:"name"`
	Email    string `json:"email" validate:"required" form:"email"`
	Phone    string `json:"phone" validate:"required" form:"phone"`
	Password string `json:"password" validate:"required" form:"password"`
}

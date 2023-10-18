package requests

type LoginRequest struct {
	Username string `json:"username" validate:"required" form:"username"`
	Password string `json:"password" validate:"required" form:"password"`
}

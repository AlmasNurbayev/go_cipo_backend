package dto

type AuthRegisterRequest struct {
	Email    string `json:"email" validate:"required,email" example:"test@test.com"`
	Password string `json:"password" validate:"required,min=4" example:"password"`
	Name     string `json:"name" validate:"omitempty" example:"test_name"`
}

type AuthRegisterResponse struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"test@test.com"`
	Password string `json:"password" validate:"required,min=4" example:"password"`
}

type AuthLoginResponse struct {
	Id int64 `json:"id"`
}

type LogoutRequest struct {
	Email     string `json:"email" validate:"required,email" example:"test@test.com"`
	SessionId string `json:"session_id" validate:"required" example:"session_id"`
}

type LogoutResponse struct {
	sessions []struct {
		SessionId string `json:"session_id" validate:"required" example:"session_id"`
		IpAddress string `json:"ip_address" validate:"required" example:"ip_address"`
	}
}

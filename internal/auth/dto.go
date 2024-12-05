package auth

type Register struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type OTPRequest struct {
	Email string `json:"email" validate:"required"`
	OTP   uint   `json:"otp" validate:"required"`
}

type ResendOTP struct {
	Email string `json:"email" validate:"required"`
}

type RegisterResponse struct {
	ID         string `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

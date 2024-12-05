package dto

type AdminRequestCreate struct {
	Name            string `form:"name" validate:"required"`
	Email           string `form:"email" validate:"required,email"`
	Password        string `form:"password" validate:"required"`
	ConfirmPassword string `form:"confirm_password" validate:"required,eqfield=Password"`
	Role            string `form:"role" validate:"required"`
}

type AdminUpdateRequest struct {
	Name        string `form:"name"`
	Email       string `form:"email" validate:"email"`
	OldPassword string `form:"old_password"`
	NewPassword string `form:"new_password"`
	Role        string `form:"role"`
}

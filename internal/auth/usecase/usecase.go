package auth

import (
	adm "github.com/sawalreverr/recything/internal/admin/repository"
	a "github.com/sawalreverr/recything/internal/auth"
	"github.com/sawalreverr/recything/internal/helper"
	u "github.com/sawalreverr/recything/internal/user"
	"github.com/sawalreverr/recything/pkg"
)

type authUsecase struct {
	userRepository  u.UserRepository
	adminRepository adm.AdminRepository
}

func NewAuthUsecase(userRepo u.UserRepository, admRepo adm.AdminRepository) a.AuthUsecase {
	return &authUsecase{userRepository: userRepo, adminRepository: admRepo}
}

func (uc *authUsecase) RegisterUser(user a.Register) (*u.User, error) {
	emailFound, _ := uc.userRepository.FindByEmail(user.Email)
	if emailFound != nil {
		return nil, pkg.ErrEmailAlreadyExists
	}

	// phoneFound, _ := uc.userRepository.FindByPhoneNumber(user.PhoneNumber)
	// if phoneFound != nil {
	// 	return nil, pkg.ErrPhoneNumberAlreadyExists
	// }

	lastID, _ := uc.userRepository.FindLastID()
	newID := helper.GenerateCustomID(lastID, "USR")

	hashedPass, _ := helper.GenerateHash(user.Password)

	newUser := u.User{
		ID:    newID,
		Name:  user.Name,
		Email: user.Email,
		// PhoneNumber: user.PhoneNumber,
		Password:   hashedPass,
		OTP:        helper.GenerateOTP(),
		IsVerified: false,
	}

	createdUser, err := uc.userRepository.Create(newUser)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	return createdUser, nil
}

func (uc *authUsecase) LoginUser(user a.Login) (string, error) {
	userFound, err := uc.userRepository.FindByEmail(user.Email)
	if err != nil {
		return "", pkg.ErrUserNotFound
	}

	ok := helper.ComparePassword(userFound.Password, user.Password)
	if !ok {
		return "", pkg.ErrPasswordInvalid
	}

	if !userFound.IsVerified {
		return "", pkg.ErrNeedToVerify
	}

	token, err := helper.GenerateTokenJWT(userFound.ID, "user")

	return token, err
}

func (uc *authUsecase) VerifyOTP(user a.OTPRequest) error {
	userFound, err := uc.userRepository.FindByEmail(user.Email)
	if err != nil {
		return pkg.ErrUserNotFound
	}

	if userFound.IsVerified {
		return pkg.ErrUserAlreadyVerified
	}

	if user.OTP != userFound.OTP {
		return pkg.ErrOTPInvalid
	}

	userFound.IsVerified = true
	if err := uc.userRepository.Update(*userFound); err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (uc *authUsecase) UpdateOTP(email string) (uint, error) {
	userFound, err := uc.userRepository.FindByEmail(email)
	if err != nil {
		return 0, pkg.ErrUserNotFound
	}

	if userFound.IsVerified {
		return 0, pkg.ErrUserAlreadyVerified
	}

	userFound.OTP = helper.GenerateOTP()

	if err := uc.userRepository.Update(*userFound); err != nil {
		return 0, pkg.ErrStatusInternalError
	}

	return userFound.OTP, nil
}

func (uc *authUsecase) LoginAdmin(admin a.Login) (string, error) {
	adminFound, err := uc.adminRepository.FindAdminByEmail(admin.Email)
	if err != nil {
		return "", pkg.ErrUserNotFound
	}

	ok := helper.ComparePassword(adminFound.Password, admin.Password)
	if !ok {
		return "", pkg.ErrPasswordInvalid
	}

	token, err := helper.GenerateTokenJWT(adminFound.ID, "admin")

	if adminFound.Role == "super admin" {
		token, err = helper.GenerateTokenJWT(adminFound.ID, "super admin")
	}

	return token, err
}

package repository

import "github.com/sawalreverr/recything/internal/admin/entity"

type AdminRepository interface {
	CreateDataAdmin(admin *entity.Admin) (*entity.Admin, error)
	UpdateDataAdmin(admin *entity.Admin, id string) (*entity.Admin, error)
	FindAdminByEmail(email string) (*entity.Admin, error)
	FindAdminByID(id string) (*entity.Admin, error)
	GetDataAllAdmin(limit int, offset int) ([]entity.Admin, int, error)
	FindLastIdAdmin() (string, error)
	DeleteAdmin(id string) error
}

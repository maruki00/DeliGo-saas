package contracts

import (
	"deligo/internal/auth/domain/entities"
	"deligo/internal/auth/infra/models"
	shared_entities "deligo/internal/shared/domain/entities"
	shared_models "deligo/internal/shared/infra/models"
)

type IAuthRepository interface {
	// Login(login, password string) (*auth_domain_dtos.AuthDTO, error)
	CheckToken(token string) bool
	Login(login, password string) (*shared_models.User, error)
	CreateAuth(token string, user *shared_models.User) *models.Auth
	ClearToken(id int)
	Register(user shared_entities.UserEntity) (shared_entities.UserEntity, error)
	CleanPins(email string)
	TwoFactoryCreate(twofactory entities.TwoFactoryPinEntity) (bool, error)
	TwoFactoryConfirm(email string, pin int) (bool, error)
	LockUser(email string, lock string)
	Logout(token string)
}

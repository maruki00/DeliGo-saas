package services

import (
	"deligo/internal/auth/domain/dtos"
	shared_contracts "deligo/internal/shared/domain/contracts"
)

type AuthService struct {
	// repo    contracts.IAuthRepository
	// outport ports.AuthOutputPort
}

func NewAuthService(
// repo contracts.IAuthRepository,
// outport ports.AuthOutputPort,
) *AuthService {
	return &AuthService{
		// repo:    repo,
		// outport: outport,
	}
}

func (obj *AuthService) Login(dto dtos.LoginDTO) shared_contracts.ViewModel {

	return nil
	// user, err := obj.repo.Login(dto.Login, utils.Md5Hash(dto.Password))
	// if err != nil {
	// 	return obj.outport.Error(shared_models.ResponseModel{
	// 		Status:  400,
	// 		Message: "Error",
	// 		Error:   err.Error(),
	// 		Result:  nil,
	// 	})
	// }

	// token, err := pkgJwt.GenerateToken(user)
	// if err != nil {
	// 	return obj.outport.Error(shared_models.ResponseModel{
	// 		Status:  400,
	// 		Message: "Error",
	// 		Error:   err.Error(),
	// 		Result:  nil,
	// 	})
	// }

	// obj.repo.ClearToken(user.GetID())
	// auth := obj.repo.CreateAuth(token, user)
	// obj.repo.LockUser(auth.Email, "1")
	// obj.repo.CleanPins(auth.Email)
	// ok, err := obj.repo.TwoFactoryCreate(&models.TwoFactoryPin{
	// 	Pin:   rand.Intn(99999999),
	// 	Email: dto.Login,
	// })

	// if !ok || err != nil {
	// 	return obj.outport.Error(shared_models.ResponseModel{
	// 		Status:  400,
	// 		Message: "Error",
	// 		Error:   err.Error(),
	// 		Result:  nil,
	// 	})
	// }
	// return obj.outport.Success(shared_models.ResponseModel{
	// 	Status:  200,
	// 	Message: "Success",
	// 	Error:   nil,
	// 	Result:  gin.H{"result": token},
	// })
}

func (obj *AuthService) Register(dto dtos.RegisterDTO) shared_contracts.ViewModel {
	return nil

	// dt := time.Now()
	// formattedTime := dt.Format("2006-01-02 15:04:05")
	// user := &shared_models.User{
	// 	UserName:  dto.UserName,
	// 	FullName:  dto.FullName,
	// 	Email:     dto.Email,
	// 	Address:   dto.Address,
	// 	Password:  utils.Md5Hash(dto.Password),
	// 	UserType:  "customer",
	// 	UserLevel: dto.UserLevel,
	// 	IsOnline:  0,
	// 	IsLocked:  1,
	// 	LastSeen:  formattedTime,
	// 	CreatedAt: formattedTime,
	// 	UpdatedAt: formattedTime,
	// }
	// _, err := obj.repo.Register(user)
	// if err != nil {
	// 	return obj.outport.Error(shared_models.ResponseModel{
	// 		Status:  400,
	// 		Message: "Error",
	// 		Error:   err.Error(),
	// 		Result:  nil,
	// 	})
	// }
	// return obj.outport.Error(shared_models.ResponseModel{
	// 	Status:  200,
	// 	Message: "Success",
	// 	Error:   nil,
	// 	Result:  nil,
	// })
}

func (obj *AuthService) TwoFactoryConfirm(dto dtos.TwoFactoryConfirmDTO) shared_contracts.ViewModel {
	return nil
	// _, err := obj.repo.TwoFactoryConfirm(dto.Email, dto.Pin)
	// if err != nil {
	// 	return obj.outport.Error(shared_models.ResponseModel{
	// 		Status:  400,
	// 		Message: "Error",
	// 		Error:   err.Error(),
	// 		Result:  nil,
	// 	})
	// }
	// obj.repo.LockUser(dto.Email, "0")
	// return obj.outport.Success(shared_models.ResponseModel{
	// 	Status:  200,
	// 	Message: "Success",
	// 	Error:   nil,
	// 	Result:  nil,
	// })
}

func (obj *AuthService) Logout(dto dtos.LogoutDTO) {
	//obj.repo.Logout(dto.Token)
}

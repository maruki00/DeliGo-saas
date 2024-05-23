package Authusecases

import (
	"delivery/Services/Auth/Domain/DTOs"
	user_repositories "delivery/Services/Auth/Infrastructure/Repositories"
	"fmt"
)

type RegisteruserCase struct {
	repo *user_repositories.AuthRepository
}

// userName string, fullName string, email string, password string, country string, city string, street string, house int, flat int
func (obj *LoginuserCase) Register(dto DTOs.UserDTO) string {

	fmt.Println("Before : ", dto.UserName,
		dto.FullName,
		dto.Email,
		dto.Password,
		dto.Country,
		dto.City,
		dto.Street,
		dto.House,
		dto.Flat)
	fmt.Println(dto)
	_ = obj.repo.Register(
		dto.UserName,
		dto.FullName,
		dto.Email,
		dto.Password,
		dto.Country,
		dto.City,
		dto.Street,
		dto.House,
		dto.Flat,
	)

	// fmt.Print("Hello world")
	// if err != nil {
	// 	return "invalid Data"
	// }
	return "success"
}

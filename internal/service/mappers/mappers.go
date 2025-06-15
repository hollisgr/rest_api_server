package mappers

import "rest_api/internal/service/dto"

func FromCreateUserToDB(u dto.UserCreateDTO) dto.DBUserCreateDTO {
	res := dto.DBUserCreateDTO{
		Login:        u.Login,
		FirstName:    u.FirstName,
		SecondName:   u.SecondName,
		Email:        u.Email,
		PasswordHash: "",
	}
	return res
}

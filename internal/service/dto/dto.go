package dto

type UserAuthDTO struct {
	Login    string `json:"login" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required"`
}

type UserCreateDTO struct {
	Login      string `json:"login" validate:"required,min=5,max=20"`
	FirstName  string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	Email      string `json:"email" validate:"required,email,contains=@"`
	Password   string `json:"password" validate:"required"`
}

type DBUserCreateDTO struct {
	Login        string `json:"login" validate:"required,min=5,max=20"`
	FirstName    string `json:"first_name" validate:"required"`
	SecondName   string `json:"second_name" validate:"required"`
	Email        string `json:"email" validate:"required,email,contains=@"`
	PasswordHash string `json:"password" validate:"required"`
}

type UserLoadDTO struct {
	Id           int    `json:"id" validate:"gte=0"`
	Login        string `json:"login" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	SecondName   string `json:"second_name" validate:"required"`
	Email        string `json:"email" validate:"required,email,contains=@"`
	PasswordHash string `json:"password_hash" validate:"required"`
}

type UserListLoadDTO struct {
	Id         int    `json:"id" validate:"gte=0"`
	Login      string `json:"login" validate:"required,min=5,max=20"`
	FirstName  string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	Email      string `json:"email" validate:"required,email,contains=@"`
}

type UserUpdateDTO struct {
	Id         int    `json:"id" validate:"gte=0"`
	Login      string `json:"login" validate:"required,min=5,max=20"`
	FirstName  string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	Email      string `json:"email" validate:"required,email,contains=@"`
}

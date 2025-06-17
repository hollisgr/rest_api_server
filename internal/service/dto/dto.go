package dto

// WEB USERS

type WebUserAuth struct {
	Login    string `json:"login" validate:"required,min=5,max=20" example:"login"`
	Password string `json:"password" validate:"required,min=5,max=20" example:"password"`
}

type WebUserCreate struct {
	Login      string `json:"login" validate:"required,min=5,max=20" example:"login"`
	FirstName  string `json:"first_name" validate:"required,min=5,max=20" example:"first_name"`
	SecondName string `json:"second_name" validate:"required,min=5,max=20" example:"second_name"`
	Email      string `json:"email" validate:"required,email,contains=@" example:"example@email.com"`
	Password   string `json:"password" validate:"required,min=5,max=20" example:"password"`
}

type WebUserLoad struct {
	Id         int    `json:"id" validate:"gte=0"`
	Login      string `json:"login" validate:"required,min=5,max=20" example:"login"`
	FirstName  string `json:"first_name" validate:"required,min=5,max=20" example:"first_name"`
	SecondName string `json:"second_name" validate:"required,min=5,max=20" example:"second_name"`
	Email      string `json:"email" validate:"required,email,contains=@" example:"example@email.com"`
}

type WebUserListLoad struct {
	Id         int    `json:"id" validate:"gte=0"`
	Login      string `json:"login" validate:"required,min=5,max=20" example:"login"`
	FirstName  string `json:"first_name" validate:"required,min=5,max=20" example:"first_name"`
	SecondName string `json:"second_name" validate:"required,min=5,max=20" example:"second_name"`
	Email      string `json:"email" validate:"required,email,contains=@" example:"example@email.com"`
}

type WebUserUpdate struct {
	Id         int    `json:"id" validate:"gte=0"`
	Login      string `json:"login" validate:"required,min=5,max=20" example:"login"`
	FirstName  string `json:"first_name" validate:"required,min=5,max=20" example:"first_name"`
	SecondName string `json:"second_name" validate:"required,min=5,max=20" example:"second_name"`
	Email      string `json:"email" validate:"required,email,contains=@" example:"example@email.com"`
}

// DBUSERS

type DBUserCreate struct {
	Login        string `json:"login" validate:"required,min=5,max=20"`
	FirstName    string `json:"first_name" validate:"required"`
	SecondName   string `json:"second_name" validate:"required"`
	Email        string `json:"email" validate:"required,email,contains=@"`
	PasswordHash string `json:"password_hash"`
}

type DBUserLoad struct {
	Id           int    `json:"id" validate:"gte=0"`
	Login        string `json:"login" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	SecondName   string `json:"second_name" validate:"required"`
	Email        string `json:"email" validate:"required,email,contains=@"`
	PasswordHash string `json:"password_hash"`
}

type DBUserListLoad struct {
	Id           int    `json:"id" validate:"gte=0"`
	Login        string `json:"login" validate:"required,min=5,max=20"`
	FirstName    string `json:"first_name" validate:"required"`
	SecondName   string `json:"second_name" validate:"required"`
	Email        string `json:"email" validate:"required,email,contains=@"`
	PasswordHash string `json:"password_hash"`
}

type DBUserUpdate struct {
	Id         int    `json:"id" validate:"gte=0"`
	Login      string `json:"login" validate:"required,min=5,max=20"`
	FirstName  string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	Email      string `json:"email" validate:"required,email,contains=@"`
}

// JWT TOKEN CREATE

type JWTTokenCreate struct {
	Id         int    `json:"id" validate:"gte=0"`
	Login      string `json:"login" validate:"required,min=5,max=20"`
	FirstName  string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	Email      string `json:"email" validate:"required,email,contains=@"`
}

package dto

// WEB USERS

type WebUserAuth struct {
	Login    string `json:"login" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required"`
}

type WebUserCreate struct {
	Login      string `json:"login" validate:"required,min=5,max=20"`
	FirstName  string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	Email      string `json:"email" validate:"required,email,contains=@"`
	Password   string `json:"password" validate:"required"`
}

type WebUserLoad struct {
	Id         int    `json:"id" validate:"gte=0"`
	Login      string `json:"login" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	Email      string `json:"email" validate:"required,email,contains=@"`
}

type WebUserListLoad struct {
	Id         int    `json:"id" validate:"gte=0"`
	Login      string `json:"login" validate:"required,min=5,max=20"`
	FirstName  string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	Email      string `json:"email" validate:"required,email,contains=@"`
}

type WebUserUpdate struct {
	Id         int    `json:"id" validate:"gte=0"`
	Login      string `json:"login" validate:"required,min=5,max=20"`
	FirstName  string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	Email      string `json:"email" validate:"required,email,contains=@"`
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

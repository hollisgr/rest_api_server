package user

type User struct {
	Id           string `json:"id"`
	Login        string `json:"login"`
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

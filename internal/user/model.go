package user

// capital chars - exportable, lowercase - non-exportable

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
	Password     string `json:"password"`
	Email        string `json:"email"`
}

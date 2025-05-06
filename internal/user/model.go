package user

// capital chars - exportable, lowercase - non-exportable

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	TG_ID        int64  `json:"telegram_id"`
}

type TGAdmin struct {
	TG_ID     int64  `json:"telegram_id"`
	ADMIN_PWD string `json:"admin_password"`
}

type TGUser struct {
	ID          int    `json:"id"`
	TG_ID       int64  `json:"telegram_id"`
	TG_USERNAME string `json:"tg_username"`
	Role_id     int    `json:"role_id"`
}

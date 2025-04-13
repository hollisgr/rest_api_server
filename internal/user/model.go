package user

// capital chars - exportable, lowercase - non-exportable

type User struct {
	UID          string `json:"uid" bson:"_id,omitempty"`
	ID           int64  `json:"id" bson:"id,omitempty"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"passwordHash" bson:"passwordHash"`
	Password     string `json:"password" bson:"-"`
	Email        string `json:"email" bson:"email"`
}

package user

import (
	"context"
	"fmt"
)

func isAdmin(storage Storage, u *User) error {

	adminRole := 2

	tguser, err := storage.FindTgUser(context.Background(), u.TG_ID)

	if err != nil {
		return fmt.Errorf("tg_user not found")
	}

	if tguser.Role_id != adminRole {
		return fmt.Errorf("access not granted")
	}

	return nil
}

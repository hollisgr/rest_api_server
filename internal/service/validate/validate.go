package validate

import "fmt"

const (
	OK int = iota
	SHORTPWD
	LONGPWD
	FORBIDDENCHARS
)

func validatePassword(password string) int {
	len := len(password)
	code := OK
	if len < 5 {
		code = SHORTPWD
	}
	if len > 20 {
		code = LONGPWD
	} else {
		runesArr := []rune(password)
		for i := 0; i < len; i++ {
			if runesArr[i] < '!' || runesArr[i] > '~' {
				code = FORBIDDENCHARS
			}
		}
	}
	return code
}

func CheckPassword(password string) (err error) {
	code := validatePassword(password)

	if code != OK {
		switch code {
		case LONGPWD:
			err = fmt.Errorf("password max length 20")
		case SHORTPWD:
			err = fmt.Errorf("password min length 5")
		case FORBIDDENCHARS:
			err = fmt.Errorf("password contains forbidden chars")
		}
	}
	return err
}

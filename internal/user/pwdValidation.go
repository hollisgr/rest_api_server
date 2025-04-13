package user

import (
	"fmt"
	"net/http"
)

const (
	OK int = iota
	SHORTPWD
	LONGPWD
	FORBIDDENCHARS
)

func isValid(pwd string) int {
	len := len(pwd)
	code := OK
	if len < 5 {
		code = SHORTPWD
	}
	if len > 20 {
		code = LONGPWD
	} else {
		runesArr := []rune(pwd)
		for i := 0; i < len; i++ {
			if runesArr[i] < '!' || runesArr[i] > '~' {
				code = FORBIDDENCHARS
			}
		}
	}
	return code
}

func pwdValidation(w http.ResponseWriter, h *handler, pwd string) error {
	val := isValid(pwd)
	var err error
	switch val {
	case SHORTPWD:
		h.logger.Infoln("pwd is short")
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Password is too short")
		err = fmt.Errorf("pwd is short")
		break
	case LONGPWD:
		h.logger.Infoln("pwd is long")
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Password is too long")
		err = fmt.Errorf("pwd is long")
		break
	case FORBIDDENCHARS:
		h.logger.Infoln("pwd contains forbidden chars")
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Password contains forbidden chars")
		err = fmt.Errorf("pwd contains forbiddent chars")
		break
	default:
		err = nil
		break
	}
	return err
}

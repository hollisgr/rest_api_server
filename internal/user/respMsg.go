package user

import "encoding/json"

type RestMsg struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (r *RestMsg) NewError(status int, code string, msg string) *RestMsg {
	return &RestMsg{
		Status:  status,
		Code:    code,
		Message: msg,
	}
}

func (r *RestMsg) CreateMsgJson(status int, code string, msg string) []byte {
	err := r.NewError(status, code, msg)
	result, _ := json.Marshal(err)
	return result
}

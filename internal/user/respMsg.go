package user

import (
	"encoding/json"
	"net/http"
)

type RestMsg struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (r *RestMsg) NewMsg(status int, code string, msg string) *RestMsg {
	return &RestMsg{
		Status:  status,
		Code:    code,
		Message: msg,
	}
}

func (r *RestMsg) SendMsgJson(w http.ResponseWriter, status int, code string, msg string) []byte {
	NewMsg := r.NewMsg(status, code, msg)
	result, _ := json.Marshal(NewMsg)
	w.WriteHeader(status)
	w.Write(result)
	return result
}

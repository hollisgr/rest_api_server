package user

import "encoding/json"

type h_error struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (h *h_error) NewError(status int, code string, msg string) *h_error {
	return &h_error{
		Status:  status,
		Code:    code,
		Message: msg,
	}
}

func (h *h_error) CreateErrorJson(status int, code string, msg string) []byte {
	err := h.NewError(status, code, msg)
	result, _ := json.Marshal(err)
	return result
}

package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, status int, msg string, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := APIResponse{
		Status: "ok",
		Msg:    msg,
		Code:   code,
		Data:   data,
	}

	json.NewEncoder(w).Encode(res)
}

func Error(w http.ResponseWriter, status int, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := APIResponse{
		Status: "error",
		Msg:    msg,
		Code:   code,
	}

	json.NewEncoder(w).Encode(res)
}

package util

import (
	"encoding/json"
	"net/http"
)

type MyResult struct {
	Error   *MyError    `json:"error"`
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

type MyError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func WriteSuccess(w http.ResponseWriter, result interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body, err := json.Marshal(result)
	if err != nil {
		return err
	}
	w.Write(body)
	return nil
}

func WriteFailed(w http.ResponseWriter, code int, message string) {

}

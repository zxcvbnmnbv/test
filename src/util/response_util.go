package util

import (
	"encoding/json"
	"fmt"
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
	var resp MyResult
	resp.Success = true
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp.Result = result
	body, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	w.Write(body)

	return nil
}

func WriteFailed(w http.ResponseWriter, code int64, message string) {
	var resp = MyResult{
		Error: &MyError{
			Code:    code,
			Message: message,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header()
	body, err := json.Marshal(resp)
	if err != nil {
		// code_review_change: 错误处理原则：
		// 1. 错误一定高输出日志
		// 2. 当前函数能处理的就在当前处理，当前无法处理的就要抛出去
		fmt.Println("resp marshal fail in [WriteFailed]")
		return
	}
	w.Write(body)
}

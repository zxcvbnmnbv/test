package util

import (
	"fmt"
	"net/http"
	"time"
)

func HandlePanic(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("internal error: ", p)
			}
		}()
		fmt.Printf("%s\n 接口%s被调用了\n", time.Now(), request.RequestURI)
		f(writer, request)
	}
}

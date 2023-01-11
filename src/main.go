package main

import (
	"example.com/m/v2/src/db"
	"example.com/m/v2/src/interfunc"
	"example.com/m/v2/src/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	err := db.InitDB()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	defer DB.Close()
	r := mux.NewRouter()
	/*
		code_review_change:
			1. {v1-9}/{服务名}/{服务域}/{接口用途}
	*/

	/*
		panic和recover
	*/
	r.HandleFunc("/v1/usersrv/user_info/query_by_id", interfunc.QueryUserInfoByUserID).Methods(http.MethodGet)
	r.HandleFunc("/v2/usersrv/user_info/query_by_id", util.HandlePanic(interfunc.QueryUserInfoByUserID)).Methods(http.MethodGet)
	r.HandleFunc("/v1/usersrv/user_info/add_user_info", interfunc.AddUserInfo).Methods(http.MethodPost)
	r.HandleFunc("/v1/usersrv/user_info/delete_by_id", interfunc.DeleteUserInfoByUserId).Methods(http.MethodPost)
	r.HandleFunc("/v1/usersrv/user_info/change_by_id", util.HandlePanic(interfunc.ChangeUserInfoByUserId)).Methods(http.MethodPost)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

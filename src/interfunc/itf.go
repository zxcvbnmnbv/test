package interfunc

import (
	"encoding/json"
	"errors"
	"example.com/m/v2/src/db"
	"example.com/m/v2/src/util"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

/*
		code_review_change:
	 	1. 函数命名，见名知意
		2. 关键流程，日志记录
*/
// QueryUserInfoByUserID 通过用户id查询对应的用户信息
func QueryUserInfoByUserID(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.WriteFailed(w, 100000, "id格式转换失败")
		return
	}
	fmt.Println("接受到查询请求,user_id:", idStr)
	user, err := db.QueryUserInfoByUserID(int64(id))
	if err != nil {
		util.WriteFailed(w, 100001, "用户id不存在")
		return
	}
	util.WriteSuccess(w, user)
	fmt.Println("请求完成，查询用户信息成功")
}

// DeleteUserInfoByUserId 通过用户id删除对应的用户信息
func DeleteUserInfoByUserId(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		util.WriteFailed(w, 100002, "get request body message fail")
		return
	}
	fmt.Println("接受到删除请求,user_id:", user.Id)
	err = db.DeleteUserInfoByUserId(user)
	if err != nil {
		util.WriteFailed(w, 100003, "userId is not exit,delete fail")
		return
	}
	util.WriteSuccess(w, "delete success")
	fmt.Println("请求完成，删除用户信息成功")
}

// AddUserInfo 添加用户的基本信息
func AddUserInfo(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		util.WriteFailed(w, 100002, "get request body message fail")
		return
	}
	err = db.InsertUserInfo(user)
	if err != nil {
		util.WriteFailed(w, 100004, "addUseInfo fail")
		return
	}
	util.WriteSuccess(w, "addUseInfo success")
	fmt.Println("请求完成，添加用户信息成功")
}

// ChangeUserInfoByUserId 通过用户id修改对应的用户信息
func ChangeUserInfoByUserId(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		util.WriteFailed(w, 100002, "get request body message fail")
		return
	}
	fmt.Println("接受到修改请求,user_id:", user.Id)
	err = db.ChangeUserInfoByUserId(user)
	if err != nil {
		util.WriteFailed(w, 100005, "change fail")
		return
	}
	util.WriteSuccess(w, "change sucess")
	fmt.Println("请求完成，修改用户信息成功")
}

// DecodeUser 从http的post请求中解码出用户的body信息
func DecodeUser(r *http.Request) (*db.User, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("read error")
	}
	var user = db.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &user, nil
}

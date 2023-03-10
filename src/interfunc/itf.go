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
	//id, err := strconv.Atoi(idStr)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// code_review_change:
		util.WriteFailed(w, util.ERR_WRONG_PARAM, "id格式转换失败")
		return
	}
	var u *db.User
	u.Id = 10
	fmt.Println("接受到查询请求,user_id:", idStr)
	user, err := db.QueryUserInfoByUserID(id)
	if err != nil {
		util.WriteFailed(w, util.ERR_USER_NOT_EXIST, "用户id不存在")
		return
	}
	fmt.Printf("请求完成，查询用户信息成功,用户信息：%+v", user)
	util.WriteSuccess(w, user)
}

// DeleteUserInfoByUserId 通过用户id删除对应的用户信息
func DeleteUserInfoByUserId(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		util.WriteFailed(w, util.ERR_REQU_BODY_MESS, "get request body message fail")
		return
	}
	fmt.Println("接受到删除请求,user_id:", user.Id)
	err = db.DeleteUserInfoByUserId(user)
	if err != nil {
		util.WriteFailed(w, util.ERR_USER_NOT_EXIST, "userId is not exit,delete fail")
		return
	}
	util.WriteSuccess(w, "delete success")
	fmt.Println("请求完成，删除用户信息成功")
}

// AddUserInfo 添加用户的基本信息
func AddUserInfo(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		util.WriteFailed(w, util.ERR_REQU_BODY_MESS, "get request body message fail")
		return
	}
	fmt.Printf("接收到添加用户请求，接收用户信息成功,用户信息为：%+v", user)
	err = db.InsertUserInfo(user)
	if err != nil {
		util.WriteFailed(w, util.ERR_USER_INFO, "addUseInfo fail")
		return
	}
	util.WriteSuccess(w, "addUseInfo success")
	fmt.Printf("请求完成，添加用户信息成功，用户信息为：%+v", user)
}

// ChangeUserInfoByUserId 通过用户id修改对应的用户信息
func ChangeUserInfoByUserId(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		util.WriteFailed(w, util.ERR_REQU_BODY_MESS, "get request body message fail")
		return
	}
	fmt.Println("接受到修改请求,user_id:", user.Id)
	err = db.ChangeUserInfoByUserId(user)
	if err != nil {
		util.WriteFailed(w, util.ERR_USER_NOT_EXIST, "change fail")
		return
	}
	util.WriteSuccess(w, "change sucess")
	fmt.Printf("请求完成，修改用户信息成功，修改后的用户信息为：%+v", user)
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

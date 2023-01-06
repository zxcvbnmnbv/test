package interfunc

import (
	"encoding/json"
	"errors"
	"example.com/m/v2/src/db"
	"example.com/m/v2/src/util"
	"io"
	"net/http"
	"strconv"
)

func Query(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		io.WriteString(w, "query fail")
		return
	}
	user, err := db.QueryUserInfoByUserID(int64(id))
	if err != nil {
		io.WriteString(w, "query fail")
		return
	}
	util.WriteSuccess(w, user)

}
func Delete(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		io.WriteString(w, "delete fail")
		return
	}
	err = db.DeleteUserInfoByUserId(user)
	if err != nil {
		io.WriteString(w, "delete fail")
		return
	}
	util.WriteSuccess(w, "delete success")
}

func Add(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		io.WriteString(w, "mistake")
		return
	}
	err = db.InsertUserInfo(user)
	if err != nil {
		io.WriteString(w, "addUseInfo fail")
		return
	}
	util.WriteSuccess(w, "addUseInfo success")

}
func Change(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		io.WriteString(w, "change fail")
		return
	}
	err = db.ChangeUserInfoByUserId(user)
	if err != nil {
		io.WriteString(w, "change fail")
		return
	}
	util.WriteSuccess(w, "change sucess")
}

func DecodeUser(r *http.Request) (*db.User, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("read error!")
	}
	var user = db.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &user, nil
}

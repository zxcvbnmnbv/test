package db

import (
	"encoding/json"
	"testing"
)

func TestInsertUser(t *testing.T) {
}

func TestJsonMarshal(t *testing.T) {
	var user = User{
		Id:       1,
		Name:     "张三",
		Password: "123",
		Status:   1,
	}

	body, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(body))
}
func TestJsonUnMarshal(t *testing.T) {
	var userStr = `{"id":1,"name":"张三","password":"123","status":1,"is_new":false}`
	var user User
	err := json.Unmarshal([]byte(userStr), &user)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", user)
	t.Logf("name:%s", user.Name)
}

package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "mydb"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   int8   `json:"status"`
}

var DB *sql.DB

// InitDB 连接数据库
func InitDB() error {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	var err error
	DB, err = sql.Open("mysql", path)
	if err != nil {
		fmt.Println("数据库初始化失败")
		return errors.New("InitDB" + err.Error())
	}
	DB.SetConnMaxLifetime(100) //设置数据库最大连接数
	DB.SetMaxIdleConns(10)     //设置上数据库最大闲置连接数
	//验证连接
	err = DB.Ping()
	if err != nil {
		fmt.Println("open database fail")
		return errors.New("InitDB" + err.Error())
	}
	fmt.Println("connnect success")
	return nil
}

// InsertUserInfo 增加用户信息到数据库中
func InsertUserInfo(user *User) error {
	// code_review_change: 当前场景无需开启事务

	//将参数传递到sql语句中并且执行
	if DB == nil {
		fmt.Println("DB 是个空指针")
		return errors.New("[InsertUserInfo] ")
	}
	_, err := DB.Exec("INSERT INTO user (`name`, `password`) VALUES (?, ?)", user.Name, user.Password)
	if err != nil {
		fmt.Println("Mysql InsertUserInfo Exec fail")
		return errors.New("[InsertUserInfo] " + err.Error())
	}
	return nil
}

// DeleteUserInfoByUserId 通过用户ID删除用户信息
func DeleteUserInfoByUserId(user *User) error {
	//设置参数以及执行sql语句
	if DB == nil {
		fmt.Println("DB 是个空指针")
		return errors.New("[DeleteUserInfoByUserId] ")
	}
	if user.Id <= 0 {
		fmt.Println("用户ID不存在")
		return errors.New("[DeleteUserInfoByUserId] ")
	}
	_, err := DB.Exec("DELETE FROM user WHERE id = ?", user.Id)
	if err != nil {
		fmt.Println("Mysql DeleteUserInfoByUserId Exec fail")
		return errors.New("DeleteUser" + err.Error())
	}
	// code_review_change: error处理

	return nil
}

// ChangeUserInfoByUserId 更改用户id对应的用户信息
func ChangeUserInfoByUserId(user *User) error {
	//设置参数以及执行sql语句
	if DB == nil {
		fmt.Println("DB 是个空指针")
		return errors.New("ChangeUserInfoByUserId ")
	}
	if user.Id <= 0 {
		fmt.Println("用户ID不存在")
		return errors.New("ChangeUserInfoByUserId ")
	}
	_, err := DB.Exec("UPDATE user SET name=?,password=? WHERE id = ?", user.Name, user.Password, user.Id)
	if err != nil {
		fmt.Println("Exec fail")
		return errors.New("ChangeUser" + err.Error())
	}

	return nil
}

// QueryUserInfoByUserID 根据用户ID查询用户信息
func QueryUserInfoByUserID(userID int64) (*User, error) {
	// code_review_change:
	// 1. 前置的代码判断避免无效代码的执行
	// 2. 指针类型的引用记得做空指针判断
	//		1. 空指针会引起panic，panic如果没有recover服务会挂掉
	//		2. 不要相信第三方的数据
	// if userID <= 0
	// if DB == nil
	if DB == nil {
		fmt.Println("DB 是个空指针")
		return nil, errors.New("QueryUserInfoByUserID")
	}
	if userID <= 0 {
		fmt.Println("用户ID不存在")
		return nil, errors.New("QueryUserInfoByUserID")
	}
	row := DB.QueryRow("select id, name, password, status from user where id=?", userID)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Status)
	if err != nil {
		fmt.Println("query fail")
		return nil, err
	}
	return &user, nil
}

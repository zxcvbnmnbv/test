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

// InitDB 连接数据库，注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	var err error
	DB, err = sql.Open("mysql", path)
	if err != nil {
		fmt.Println(err.Error())
	}
	DB.SetConnMaxLifetime(100) //设置数据库最大连接数
	DB.SetMaxIdleConns(10)     //设置上数据库最大闲置连接数
	//验证连接
	err = DB.Ping()
	if err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

// InsertUserInfo 增加用户信息到数据库中
func InsertUserInfo(user *User) error {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return errors.New("InsertUser" + err.Error())
	}
	defer tx.Rollback()

	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`name`, `password`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return errors.New("InsertUser" + err.Error())
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.Name, user.Password)
	if err != nil {
		fmt.Println("Exec fail")
		return errors.New("[InsertUser] " + err.Error())
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return nil
}

// DeleteUserInfoByUserId 通过用户ID删除用户信息
func DeleteUserInfoByUserId(user *User) error {

	//准备sql语句
	stmt, err := DB.Prepare("DELETE FROM user WHERE Id = ?")
	if err != nil {
		fmt.Println("delete Prepare fail")
		return errors.New("DeleteUser" + err.Error())
	}
	defer stmt.Close()
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.Id)
	if err != nil {
		fmt.Println("Exec fail")
		return errors.New("DeleteUser" + err.Error())
	}

	fmt.Println(res.RowsAffected())
	return nil
}

// ChangeUserInfoByUserId 更改用户id对应的用户信息
func ChangeUserInfoByUserId(user *User) error {

	//准备sql语句
	stmt, err := DB.Prepare("UPDATE user SET (`name`, `password`) VALUES (?, ?) WHERE Id = ?")
	if err != nil {
		fmt.Println("change Prepare fail")
		return errors.New("ChangeUser" + err.Error())
	}
	defer stmt.Close()
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.Id)
	if err != nil {
		fmt.Println("Exec fail")
		return errors.New("ChangeUser" + err.Error())
	}

	fmt.Println(res.RowsAffected())
	return nil
}

// QueryUserInfoByUserID 根据用户ID查询用户信息
func QueryUserInfoByUserID(userID int64) (*User, error) {
	row := DB.QueryRow("select id, name, password, status from user where id=?", userID)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Status)
	if err != nil {
		fmt.Println("query fail")
		return nil, err
	}
	return &user, nil
}

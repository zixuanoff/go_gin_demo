package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "bf"
)

var DB *sql.DB //长连接

type User struct {
	User_id       int64  `json:"id"`
	User_name     string `json:"username"`
	User_password string `json:"password"`
	Is_admin      bool   `json:"isAdmin"`
}

func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

func InsertUser(user User) bool {

	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}

	stmt, err := tx.Prepare("INSERT INTO user (`user_name`, `user_password`, `is_admin`) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}

	res, err := stmt.Exec(user.User_name, user.User_password, user.Is_admin)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}

	tx.Commit()
	fmt.Println("increment ID: ")
	fmt.Println(res.LastInsertId())
	return true
}

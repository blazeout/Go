package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:admin@tcp(127.0.0.1:3306)/goday16"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connet with database failed", err)
		return err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	Id   int
	Name string
	Age  int
}

func main(){
	err := initDB()
	if err != nil {
		fmt.Println("初始化失败",err)
		return
	}
	sqlStr := "select id, name, age from user where id = 1"
	var u user
	err = db.Get(&u, sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n",u)
	userList := make([]user,0,10)
	sqlStr2 := "select id, name, age from user where id > 1"
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n",userList)
}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// database/sql
// Go实现mysql_transaction示例

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	// 链接数据库
	dsc := "root:admin@tcp(127.0.0.1:3306)/goday16"
	// mysql 需要使用全部的db,不能使用 := 声明一个局部的变量
	db, err = sql.Open("mysql", dsc)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5) // 设置数据库最大空闲连接数
	return
}

type user struct {
	id int
	name string
	age int
}

func transactionDemo(){
	// 1.开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed , err : %v\n",err)
		return
	}
	// 执行多个SQL操作
	sqlStr1 := "update user set age = age-2 where id = 1"
	sqlStr2 := "update user set age = age+2 where id = 6"
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 要回滚 回滚就是将之前的操作取消改为原样
		tx.Rollback()
		return
	}
	// 执行sql1
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 要回滚
		err := tx.Rollback()
		if err != nil {
			return 
		}
		return
	}
	// 执行sql2
	// 上面两步sql都执行成功,就提交本次事务
	err = tx.Commit()
	if err != nil {
		fmt.Println("提交事务出错了要回滚")
		return
	}
	fmt.Println("事务执行成功")
}

func main()  {
	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("连接数据库成功")
	transactionDemo()
}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// database/sql
// 原生支持连接池,是并发安全的
// Go链接mysql示例

func query(){

}

func insert(){

}


func main(){
	// 链接数据库
	dsc := "root:admin@tcp(127.0.0.1:3306)/goday16"
	db, err := sql.Open("mysql", dsc)
	if err != nil {
		fmt.Printf("open %s failed , err : %v\n\n", dsc, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("连接失败 err : %v\n", err)
		return
	}
	fmt.Println("连接成功")
}
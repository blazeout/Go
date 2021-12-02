package main

import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
)

// database/sql
// 原生支持连接池,是并发安全的
// Go链接mysql示例

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

// 查询单个记录
func queryOne(id int){
	var u1 user
	// 1.写查询单条记录的sql语句
	sqlStr := "select id, name, age from user where id = ?; "
	// 2. 执行 两种方法
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	// 从连接池里面拿出一个去链接数据库查询单条记录
	// 3.拿到结果
	// 必须对rowObj对象调用scan方法,因为改方法会释放对数据库的连接, db.QueryOne返回一个Row对象,直接.调用scan方法
	fmt.Printf("%v\n",u1)
}

func queryMore(n int) {
	// 1. SQL语句
	sqlStr := "select id, name, age from user where id >= ?"
	// 2. 执行
	rows, err := db.Query(sqlStr,n)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 3. 一定记得关闭rows的连接归还给数据库连接池
	defer rows.Close()
	// 4 循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id,&u1.name, &u1.age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v\n",u1)
	}
}

func insert(){
	// 1.sql语句
	sqlStr := "insert into user(name, age) values (\"帅嘉豪\",22)"
	// 2.exec执行
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 如果是插入数据的操作,能够拿到插入数据的id值
	id, err := ret.LastInsertId()
	if err != nil {
		return
	}
	fmt.Println(id)
}

func deleteRow(deleteId int){
	// 1. sql语句
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, deleteId)
	if err != nil {
		fmt.Println(err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		return
	}
	fmt.Println("受影响的行数",affected)
}

// updataRow
func updataRow(newAge, id int) {
	sqlStr := `update user set age = ? where id > ?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	n , err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("受影响的行数",n)
}

// mysql 预处理插入多条数据

func prepareInsert(){
	sqlStr := "insert into user(name, age) values(?, ?)"
	stmt, err := db.Prepare(sqlStr) // 把sql语句先发给mysql预处理一下
	if err != nil {
		fmt.Println(err)
		return
	}
	// 后续只需要拿到这个stmt执行这个操作
	defer stmt.Close()
	var m1 = map[string]int {
		"刘想起" : 30,
		"王祥吉" : 10,
		"田硕"   : 22,
		"崔经纬" : 23,
	}
	for key, value := range m1 {
		_, err := stmt.Exec(key,value)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main(){
	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("连接数据库成功")
	//queryOne(1)
	//queryMore(3)
	//insert()
	//updataRow(99,3)
	//deleteRow(6)
	prepareInsert()
}

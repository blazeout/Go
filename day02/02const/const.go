package main

import "fmt"

const pi1 = 3.1415926
const (
	statusOK = 200
	NOTFOUND = 404
)
//批量声明常量时,如果某一行声明后没有赋值,默认和上面一行一样
const(
	n1 = 100
	n2
	n3
)
// 02const iota
// iota:类似于枚举
const(
	a1 = iota
	a2
	a3
	a4
)
const (
	b1 = iota
	b2
	_
	b3
)
const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4 = 100
)
const(
	d1,d2 = iota + 1,iota + 2 // d1 = 1 d2 = 2
	d3,d4 = iota + 1,iota + 2 // d3 = 2 d4 = 3
)
// 定义数量级
const(
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
func main(){
	const pi =  3.1415927
	//06fmt.Println(n1)
	//06fmt.Println(n2)
	//06fmt.Println(n3)
	//06fmt.Println(a1)
	//06fmt.Println(a2)
	//06fmt.Println(a3)
	//06fmt.Println(b1)
	//06fmt.Println(b2)
	//06fmt.Println(b3)
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
}
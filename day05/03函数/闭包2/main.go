package main

import "fmt"

// 无参数 返回值 为 func(int)int类型
// 闭包是一个函数,函数有一个特点:包含了他外部作用域的一个变量
// 闭包底层原理:
// 1.函数可以作为返回值
// 2. 函数内部查找变量的顺序,先在自己内部找,找不到往外面找
func adder() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := adder()
	ret2 := ret(200)
	fmt.Println(ret2)
}

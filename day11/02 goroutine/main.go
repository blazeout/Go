package main

import (
	"fmt"
	"time"
)

//func hello(i int){
//	fmt.Println("hello",i)
//}

// 程序启动后会创建一个主goroutine执行
func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i) // 闭包 函数内部用了外部的变量
		}(i) // 开启一个单独的goroutine去执行这个函数
	}
	fmt.Println("mian")
	time.Sleep(time.Second)
	// main函数结束了 由main函数启动的goroutine也都结束了
}

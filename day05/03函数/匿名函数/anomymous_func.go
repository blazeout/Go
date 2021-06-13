package main

import "fmt"

func main() {

	ff1 := func(x,y int){
		fmt.Println(x + y)
	} // 匿名函数就是没有名字的函数 多用于函数内部
	ff1(10,20)
	// 直接调用函数 = 匿名函数后面加一个()马上调用
	func (){
		fmt.Println("hello world!")
	}()

}

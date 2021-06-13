package main

import "fmt"

func sum(x int,y int)(ret int){
	return x+y
}
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中,给他起个名字,每次用到他了就直接用函数名调用就可以了
// 使代码结构更简洁
func sub(x int,y int)(){
	// 有参数 无返回值
	fmt.Println(x + y)
}

func f0()  {
	// 没有参数没有返回值的
	fmt.Println("ff")
}

func f1()int  {
	return 1
}
// 多个返回值
func f2()(int,string){
	return 1,"沙河"
}
// 参数类型的简写
func f3(x,y int)(){

}
func f4(x string,y ... int){
	fmt.Println(x)
	fmt.Println(y)
}
func main() {
	// 函数
	// 函数的定义
	i := sum(10,20)
	fmt.Println(i)
	sub(10,21)
	ret := f1()
	fmt.Println(ret)
	f0()
	_,n := f2()
	fmt.Println(n)
	f4("沙河",2,34,1,3,4)
}

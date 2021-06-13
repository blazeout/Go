package main

import "fmt"

func fo1()int{
	fmt.Println("沙河")
	return 1
}
// 函数也可以作为参数的类型
func fo2(x func()int){
	a := x()
	fmt.Println(a)
}
func fo3(x,y int)int  {
	ret := x+y
	return ret
}
func ff(x,t int)int  {
	return x + t
}
// 函数作为返回值返回
func fo4(x func()int) func(int ,int)int {
	return ff
}
func main() {
	a := fo1
	fmt.Printf("%T\n",a)
	fo2(a)
	fo4(fo1)
	fmt.Printf("%T\n",fo4)
}

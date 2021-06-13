package main

import "fmt"
// 闭包的作用 ！！！！
func f1(f func()){
	fmt.Println("this is f1")
	f()
}
func f2(x,y int){
	fmt.Println("this is f2")
	fmt.Printf("%d",x+y)
}
func f3(f func(int ,int),x,y int)func(){
	ret := func(){
		f(x,y)
	}
	return ret
}
func main() {
	 //f1(f2)
	f1(f3(f2,1,2))
}

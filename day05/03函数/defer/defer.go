package main

import "fmt"

func deferDemo(x, y int) (sum int) {
	fmt.Println("start")
	// defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("开始")
	fmt.Println("end")
	sum = x + y
	return sum
}
func main() {
	// defer语句
	sum := deferDemo(10, 20)
	fmt.Println(sum)
}

package main

import "fmt"

// 第一步 : 先把值赋给返回值
// 第二步 : 执行defer语句后
// 第三步 : 执行RET语句
func fi1() int {
	x := 5
	defer func() {
		x++ // 修改的是x 不是返回值
	}()
	return x // 5等于返回值
}


func fi2() (x int) {
	defer func() {
		x++
	}() // x是返回值 x++ 那么就是6
	return 5
}

func fi3() (y int) {
	x := 5
	defer func() {
		x++ // x = 6
	}()
	return x // y = x = 5
}
func fi4() (x int) {
	defer func(x int) {
		x++ // 这个x是拷贝一份的 不是 外面的x
	}(x)
	return 5
}
func main() {
	fmt.Println(fi1())
	fmt.Println(fi2())
	fmt.Println(fi3())
	fmt.Println(fi4())
}

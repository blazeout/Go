package main

import "fmt"

func main(){

	// 用make函数穿件一个切片

	a := make([]int,5,10)
	// 类型,长度,容量
	fmt.Printf("a = %v len = %d cap = %d",a,len(a),cap(a))
	// 切片就一个框,框柱了一个连续的内存 只能存相同类型的值
	
}

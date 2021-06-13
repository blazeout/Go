package main

import "fmt"

// 递归 : 函数自己调用自己
//永远不要高估自己
func f(n uint64)(result uint64){
	// 计算n的阶乘
	if n == 1 {
		return 1
		}
		result = n*f(n-1)
		return result
}

func main(){
	x := f(40)
	fmt.Println(x)
}

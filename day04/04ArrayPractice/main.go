package main

import "fmt"

func main(){
	// 求所有元素的和
	var arr = [...]int{1,3,5,7,8}
	sum := 0
	for _, i2 := range arr {
		sum += i2
	}
	fmt.Printf("数组的值是%d",sum)
}
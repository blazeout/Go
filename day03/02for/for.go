package main

import "fmt"

func main(){
	for i := 0; i < 10; i++ {
		// 当 i = 5 时 跳出for循环
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")
	for i := 0; i < 10; i++{
		// for 当 i = 5时 跳过此次for 循环
		if i == 5{
			continue
		}
		fmt.Println(i)
	}

}

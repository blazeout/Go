package main

import "fmt"

func main(){

	// 找出数组中和为指定值的两个元素的下标

	var arr = [6]int{4,7,2,1,8,5}
	//  找到指定和为6的两个元素的下标
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] + arr[j] == 6{
				fmt.Println("下标是",i,"和",j)
			}
		}
	}
}

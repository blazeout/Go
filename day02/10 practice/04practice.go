package main

import "fmt"

// 使用Go语言打印九九乘法表

func main(){
	var col int = 0
	for row := 1; row <= 9; row++ {
		for col = 1;col <= row;col++{
			fmt.Printf("%d * %d = %d\t",col,row,row*col)
		}
		fmt.Println()
	}

}

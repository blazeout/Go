package main

import (
	"context"
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	// 先存0,再下一次循环无法存,那么就打印0 下次循环此时i = 2,通道内无值,所以又只能存2
	// 实现的效果就是 0,2,4,6,8
	context.Background()
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Println(x)
		case ch1 <- i:

		}
	}
}

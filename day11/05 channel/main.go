package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var b chan int // 需要指定通道中元素的类型

// 没有缓冲区的chan需要有goroutine来接收它里面的值才可以放进去
func noBuffChannel() {
	b = make(chan int) // 通道要make函数初始化才能使用
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("从通道中取出了值", x)
	}()
	b <- 10
	fmt.Println("发送10到了通道中了")
	wg.Wait()
}

// 有缓存去的chan
func buffChannel() {
	fmt.Println(b)
	b = make(chan int, 10)
	wg.Add(1)
	b <- 10
	fmt.Println("发送10到了通道中了")
	b <- 20
	fmt.Println("发送20到了通道中了")
	fmt.Println(b)
	close(b)
}

func main() {
	buffChannel()
}

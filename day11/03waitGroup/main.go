package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		ret1 := rand.Int()
		ret := rand.Intn(10)
		fmt.Println(ret, ret1)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * (time.Duration(rand.Intn(300))))
	fmt.Println(i)
}

var wg sync.WaitGroup

// 利用waitGroup 实现goroutine的同步

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	// 如何知道10个goroutine都结束了
	// 当wg的count减0时就都结束了
	wg.Wait()
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS
var wg sync.WaitGroup

func a(){
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Printf("A : %d \n",i)
	}

}

func b(){
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Printf("B : %d\n",i)
	}
}

func main(){
	runtime.GOMAXPROCS(2) // 两个干活的 i 的值太小分配的一段时间内就已经执行完了
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

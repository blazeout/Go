package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 原子操作
var wg sync.WaitGroup
var x int64
var lock sync.Mutex

func add() {
	lock.Lock()
	x++
	lock.Unlock()
	wg.Done()
}

func atomicAdd(x int64){
	atomic.AddInt64(&x,1)
	wg.Done()
}

// 一开始不使用锁并发的去加操作肯定会同时访问到x
// 2.加锁但是消耗过大
func main() {
	strat := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//go add()
		go atomicAdd(x)
	}
	wg.Wait() // 5.7385ms // 2.1342ms
	fmt.Println(x)
	fmt.Println(time.Now().Sub(strat))
}

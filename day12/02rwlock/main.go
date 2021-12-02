package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock
var x = 0
var lock sync.Mutex
var wg sync.WaitGroup
var rw sync.RWMutex

func read(){
	defer wg.Done()
	//lock.Lock()
	rw.RLock()
	fmt.Println(x)
	rw.RUnlock()
	time.Sleep(time.Millisecond)
	//lock.Unlock()

}

func write(){
	defer wg.Done()
	//lock.Lock()
	rw.Lock()
	x += 1
	rw.Unlock()
	time.Sleep(time.Millisecond * 5)
	//lock.Unlock()

}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

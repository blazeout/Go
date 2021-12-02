package main

import (
	"fmt"
	"strconv"
	"sync"
)

// go语言内置的map并发操作不安全
var m = make(map[string]int)
var lock sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

var m2 sync.Map

func main() {
	//wg := sync.WaitGroup{}
	//for i := 0; i < 21; i++ {
	//	wg.Add(1)
	//	go func(n int) {
	//		lock.Lock()
	//		key := strconv.Itoa(n)
	//		set(key, n)
	//		lock.Unlock()
	//		fmt.Printf("k=:%v,v:=%v\n", key, get(key))
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	wg := sync.WaitGroup{}
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         // 必须根据sync.Map内置的Store方法设置键值对
			value, _ := m2.Load(key) // 必须根据sync.Map内置的load方法取值
			fmt.Printf("k = %v v = %v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

}

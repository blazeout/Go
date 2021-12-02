package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan int){
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}


func f2(ch1 chan int,ch2 chan int){
	defer wg.Done()
	for  i := range ch1 {
		ch2 <- i * i
	}
	close(ch2)
}

func main(){
	wg.Add(2)
	a := make(chan int,100)
	b := make(chan int,100)
	go f1(a)
	go f2(a,b)

	for ret := range b{
		fmt.Println(ret)
	}
	wg.Wait()
}

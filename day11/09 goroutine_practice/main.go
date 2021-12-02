package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//1.使用goroutine和channel实现一个计算int64随机数各位数和的程序。
//2.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
//3.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//4.主goroutine从resultChan取出结果并打印到终端输出

type job struct {
	value int64
}

type result struct {
	*job
	sum int64
}

var jobChan = make(chan *job,100)
var resultChan = make(chan *result,100)
var wg sync.WaitGroup


func produce(job1 chan<- *job){
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Int63()
		newJob := &job{
			value : x,
		}
		job1 <- newJob
		time.Sleep(time.Second)
	}
}

func calc(job1 <-chan *job,resultChan chan<- *result){
	defer wg.Done()
	for  {
		job2 := <- job1
		x := job2.value
		sum := int64(0)
		for x > 0 {
			sum += x % 10
			x = x / 10
		}
		newSum := &result{
			sum: sum,
			job : job2,
		}
		resultChan <- newSum
	}
}

func main() {
	wg.Add(1)
	go produce(jobChan)
	for i := 0; i < 24; i++ {
		go calc(jobChan,resultChan)
	}
	wg.Add(24)
	for result1 := range resultChan {
		fmt.Printf("value : %d sum : %d\n",result1.job.value,result1.sum)
	}
	wg.Wait()
}

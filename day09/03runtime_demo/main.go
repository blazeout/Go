package main

import (
	"fmt"
	"path"
	"runtime"
)

func f2(){
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(path.Base(file))
	fmt.Println(file)
	fmt.Println(line)
}

func f1(){
	f2()
}

func main(){
	f1()

}

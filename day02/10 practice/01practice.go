package main

import "fmt"

//iota 的定义
const (
	s = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main(){
  fmt.Println(KB)
  fmt.Println(MB)
  fmt.Println(GB)
  fmt.Println(TB)
}

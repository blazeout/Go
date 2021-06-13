package main

import "fmt"

// 06fmt 占位符
func main(){
	var a = 100
	// 查看变量的类型
	fmt.Printf("%T\n",a)
	// 查看变量的值
	fmt.Printf("%v\n",a)
	// 查看变量以十进制输出的值
	fmt.Printf("%d\n",a)
	// 查看变量以八进制输出的值
	fmt.Printf("%o\n",a)
	// 查看变量以十六进制输出的值
	fmt.Printf("%x\n",a)
	// 查看变量以二进制输出的值
	fmt.Printf("%b\n",a)
	s := "hello go"
	fmt.Printf("%s",s)
}

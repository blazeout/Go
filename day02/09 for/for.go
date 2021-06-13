package main

import "fmt"

func main(){
	// for 循环基本格式
	for i := 0;i<10;i++{
		fmt.Println(i)
	}
	// for 循环变种格式1
	var i = 5
	for ; i<10;i++{
		fmt.Println(i)
	}
	// 变种2
	var x = 8
	for ;x<10;{
		fmt.Println(x)
		x++
	}
	// 无限循环
	//for{
	//	fmt.Println(x)
	//}
	// for range 循环
	s := "hello陕科大"
	for i,v := range s{
		fmt.Printf("%d %c\n",i,v)
	}
}

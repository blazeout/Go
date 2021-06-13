package main

import "fmt"

// 浮点数
func main()  {
	//math.MaxFloat32
	// 默认Go语言中float都是float64类型
	f1 := 1.234556
	fmt.Printf("%T",f1)
	f2 := float32(1.234552)
	//f1 = f2
	//float32不能直接赋值给float64
	fmt.Println(f2)
}

package main

import "fmt"

func main() {
	//n := 18
	//p := &n
	//fmt.Println(p)
	//fmt.Printf("%T",p)
	//str := "hello"
	//p1 := &str
	//fmt.Printf("%T",p1)
	//fmt.Println(*p1)
	// 空指针无法解引用
	// nil 无法解引用
	// 需要使用new函数申请一个内存地址

	var a = new(int) // 有自己的内存空间
	var a1 *int // 没有赋值那么就是指针类型的零值 nil
	fmt.Println(a1)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)
	// make 和 new 都是用申请内存的 new用于申请基本数据类型 int float string 返回的是类型的指针 基本很少用
	// make 是用来给 slice map channel 申请内存的 返回的是 这三个类型本身
	//var b map[string]int
	//b["沙河"] = 100
	//fmt.Println(b)


}

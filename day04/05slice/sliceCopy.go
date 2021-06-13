package main

import "fmt"

func main() {
	// copy函数 拷贝一份 修改原来的值 复制新的切片不会更改
	a1 := []int{1,3,5}
	a2 := a1
	var a3  = make([]int,3,3)
	copy(a3,a1)
	fmt.Println(a2)
	fmt.Println(a3)
	a1[0] = 100
	fmt.Println(a1,a2,a3)
	// Go语言中没有删除的专用方法 只能自己去实现
	a := []int{1,2,3,4,5,6,7}
	a = append(a[:2],a[3:]...)
	fmt.Println(a)
	x1 :=[...]int{1,3,5}
	s1 := x1[:]
	fmt.Println(s1,len(s1),cap(s1))
	s1 = append(s1[:1],s1[2:]...)
	fmt.Println(s1,len(s1),cap(s1))
	fmt.Println(s1,x1)
}

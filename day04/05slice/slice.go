package main

import "fmt"

func main(){
	// 切片 动态数组 底层是一个数组 做了封装
	var s1 []int
	var s2 []string
	fmt.Printf("%T",s1)
	fmt.Printf("%T",s2)

	s1 = []int{1,2,3,4}
	s3 := []string{"沙河","程总粗","傻逼i"}
	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Printf("len : %d cap : %d",len(s1),cap(s1))
	fmt.Printf("len : %d cap : %d",len(s3),cap(s3))
	fmt.Printf("len : %d cap : %d",len(s3),cap(s3))
	a1 := [...]int{1,3,4,5,6}
	a2 := a1[0:4]
	// 基于一个数组得到切片 左闭右开
	fmt.Println(a2)
	s4 := a1[1:]
	fmt.Printf("len : %d cap : %d",len(s4),cap(s4))

	// 切片的容量是从第一个切的元素开始的底层数组的长度 切片的长度是本身的长度
	// 切片指向了一个底层的数组
	// 切片的长度是元素的个数
	// 切片的容量是切片的第一个元素到底层数组的最后一个
	s5 := s4[2:3]
	fmt.Printf("len : %d cap : %d",len(s5),cap(s5))

}

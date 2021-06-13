package main

import (
	"fmt"
	"sort"
)

func main() {
	// 一开始make函数创建了一个长度为5的切片并且初始化为[0 0 0 0 0] 然后再append
	var a = make([]string,5,10)// string类型的切片,长度为5,底层数组容量为10
	for i := 0; i < 10; i++ {
		a = append(a,fmt.Sprintf("%v",i))
	}
	fmt.Println(a)
	a1 := []int{2,4,1,6,7}
	sort.Ints(a1)
	fmt.Println(a1)
	a2 := [...]int{1,3,5,7,9,11,13,15,17}
	s1 := a2[:]
	s1 = append(s1[:1],s1[2:]...)
	// 底层数组是先把1切出来然后将5赋值到3的位置,7赋值到5的位置 依次赋值,最后赋值完是
	//[1,5,7,9,11,13,15,17,17] 并且最后的17没有被覆盖所以还在原来的地方
	fmt.Println(s1)
	fmt.Println(a2)
}

package main

import "fmt"

func modify(a [3]int)  {
	a[0] = 100
}
func modify2(erwei [3][3]int){
	erwei[2][2] = 100
}
func main(){


	var a [3]int
	var b  = []int{1,2}
	var str = [...]string{"peaking","shanghai","changsha"}
	//a = b
	modify(a)
	// 不可以这样做 因为 [3]int 和 [4]int 不是同一类型
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("the type of str is %T\n",str)
	x := [...]int{1: 1, 3: 5}
	fmt.Println(x)
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
	fmt.Println("=====================")
	for index, value := range x {
		fmt.Println(index,value)
	}

	var ewer = [3][3]int{
		{11,22,0},
		{33,44,0},
		{55,66,0},
	}
	ewer[2][2] = 100
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(ewer[i][j])
		}
	}
	t := [3]*int{}
	fmt.Println(t)
}

package main

import "fmt"

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main(){
	var zhouling person
	zhouling.hobby = []string{"睡觉","篮球"}
	var p2 person
	p2.name = "理想"
	p2.gender = "男"
	fmt.Printf("type : %T value : %v",p2,p2)
}

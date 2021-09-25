package main

import (
	"fmt"
)

// 接口是一种类型,规定了变量有哪些方法

// 定义一个能叫的类型
type speaker interface {
	speak() //只要是实现了speak方法的变量都可以是speaker类型接口,方法签名

}

type cat struct {

}

type dog struct {

}

type person struct {

}

func (p person) speak(){
	fmt.Println("卧槽")
}

func (c cat) speak(){
	fmt.Println("喵喵喵")
}

func (d dog) speak(){
	fmt.Println("汪汪汪")
}

func da(x speaker){
	x.speak()
}

func main() {
	var d1 dog
	var c1 cat
	var p1 person
	da(d1)
	da(c1)
	da(p1)
	var ss speaker // 定义一个speaker变量 : ss
	ss = c1
	ss = p1
	ss = d1
	fmt.Println(ss)
}

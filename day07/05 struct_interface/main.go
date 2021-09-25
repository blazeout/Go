package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string2 string)
}

type cat struct {
	name string
	feet int8
}
// 方法使用值接收者
// 方法使用指针接受者
func (c *cat) move(){
	fmt.Println("猫冬")
}

func (c *cat) eat(something string) {
	fmt.Printf("猫吃%s",something)
}
func main() {

}

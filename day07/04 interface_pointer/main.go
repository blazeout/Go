package main

import "fmt"

type animal interface {
	move()
	eat(string)
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
	var a1 animal
	c1 := cat{"几个",4}
	c2 := &cat{"好咯",4}
	a1 = &c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)

}

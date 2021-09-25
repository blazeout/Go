package main

import "fmt"

// type 接口名 interface{
// 方法名1(参数1,参数2...)(返回值1,返回值2...)
// 方法名2(参数1,参数2...)(返回值1,返回值2...)
// ...
// }


type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c cat) move(){
	fmt.Println("猫冬")
}

func (c cat) eat(something string) {
	fmt.Printf("猫吃%s",something)
}

func (c chicken) move(){
	fmt.Println("激动!")
}

func (c chicken) eat(string2 string){
	fmt.Printf("鸡翅%s!\n", string2)
}

func main() {

	var a1 animal // 定义一个接口类型的变量

	bc := cat{ // 定义一个cat类型的变量bc
		name: "塔器",
		feet: 4,
	}
	a1 = bc
	bc.eat("鱼")
	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	a1.eat("halo")
	fmt.Printf("%T",a1)
}

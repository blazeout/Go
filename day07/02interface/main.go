package main

// 接口实例2

import "fmt"

type car interface {
	move()
}

type falali struct {
	brand string
}

type porsChe struct {
	brand string
}

func (f falali)move(){
	fmt.Printf("%s 要跑了",f.brand)
}

func (p porsChe)move(){
	fmt.Printf("%s 也要跑了",p.brand)
}

func move(car2 car){
	car2.move()
}

func main() {
	var f1  = falali{
		"法拉利",
	}
	var p1  = porsChe{
		brand: "保时捷",
	}
	move(f1)
	move(p1)
}

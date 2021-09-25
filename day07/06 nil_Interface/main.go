package main

import "fmt"

// 空接口没有必要起名字
// 所有类型都实现了空接口,也就是任意类型都能保存到空接口中 泛型编程
// 空接口作为函数参数 啥值都能传
func show(a interface{}){
	fmt.Printf("type : %T value : %v",a,a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{},16)
	m1["name"] = "周玲"
	m1["age"] = 18
	m1["married"] = true
	m1["hobby"] = [...]string{"唱跳","rap"}
	fmt.Println(m1)
	show(m1)
}

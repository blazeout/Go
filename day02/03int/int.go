package main

import "fmt"

// 整形
func main()  {
	var i1 = 101
	fmt.Printf("%d\n",i1)

	fmt.Printf("%o\n",i1)
	fmt.Printf("%b\n",i1)
	fmt.Printf("%x\n",i1)
	i2 := 077
	fmt.Printf("%d\n",i2)

	i3 := 0x100
	fmt.Printf("%d\n",i3)
	// 声明int8类型
	// 明确指定int类型否则就为int类型
	i4 := int8(9)
	fmt.Printf("%T",i4)

}
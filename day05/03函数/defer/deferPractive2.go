package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
	// 1. x := 1
	// 2. y := 2
	// 3.defer calc("AA", 1, 3)
	// 4.calc("A", 1, 2) // "A" 1 2 3
	// 5.x = 10
	// 6.defer calc("BB", 10, calc("B", 10, 2))
	// 7.calc("B", 10, 2) 打印 B 10 2 12
	// 8.y = 20
	// 9. BB 10 12 22
	// 10. A 1 3 4
}

package main

import "fmt"

func shahe() {
	fmt.Println("沙河")
}
func share1(string2 string) {
	fmt.Println("hello" + string2)
}
func sum1(x int, y int) int {
	sum := x + y
	return sum
}
func sum2(x, y int) (sum int) {
	sum = x + y
	return sum
}
func main() {
	var arr = [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	shahe()
	share1("理想")
	x := sum1(10, 20)
	fmt.Println(x)

}

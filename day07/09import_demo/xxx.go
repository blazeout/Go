package main

import (
	"fmt"
	calc "github.com/day07/08calc"
	"math"
)
var x = 10
const  pi  = 3.14
func init(){
	fmt.Println(x,math.Pi)

}

func main(){
	x := calc.Add(4,6)
	fmt.Println(x)
}

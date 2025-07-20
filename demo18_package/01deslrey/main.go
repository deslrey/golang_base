package main

import (
	"01deslrey/calc"
	"01deslrey/tools"
	"fmt"
)

func init() {
	fmt.Println("mian init...")
}

func main() {

	sum := calc.Add(10, 20)
	fmt.Printf("sum: %v\n", sum)
	sub := calc.Sub(10, 2)
	fmt.Printf("sub: %v\n", sub)

	b := tools.Mul(2, 4)
	fmt.Printf("b: %v\n", b)

}

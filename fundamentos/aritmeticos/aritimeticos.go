package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 3, 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicaçao =>", a*b)
	fmt.Println("Modulo =>", a%b) // resto da divisao

	// bitwise
	fmt.Println("E =>", a&b)  // 11 & 10 = 10
	fmt.Println("Ou =>", a|b) // 11 |10 = 11
	fmt.Println("xor=>", a^b) // 11^10 = 01

	c, d := 3.0, 2.0
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}

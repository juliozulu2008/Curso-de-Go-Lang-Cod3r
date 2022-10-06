package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	fmt.Println(" nova")
	fmt.Println("Linha")

	x := 3.141516
	//go nao aceita string e int como print tem que contornar como no exemplo
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("o valor de x é %.2f", x)

	a, b, c, d := 1, 1.9999, false, "opa"
	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}

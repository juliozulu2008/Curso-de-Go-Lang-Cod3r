package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador que vai contar o tamannho do meu array

	for i, numero := range numeros { // indece e numero
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // ignorando o indice e pegando apenas o numero
		fmt.Println(num)
	}
}

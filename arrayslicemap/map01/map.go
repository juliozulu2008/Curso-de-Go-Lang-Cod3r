package main

import "fmt"

func main() {
	// maps devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[99999999999] = "Julio"
	aprovados[88888888888] = "Kelly"
	aprovados[77777777777] = "Laura"
	aprovados[66666666666] = "Ravi"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[99999999999])

	delete(aprovados, 99999999999)
	fmt.Println(aprovados)
}

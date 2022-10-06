package main

import "fmt"

func main() {
	i := 1
	// ponteiros é uma referencia de memoria
	// Go não tem aritmética de ponteiros
	var p *int = nil

	p = &i // pegando o endereço da variavel e atribuindo ao ponteiro
	*p++   // Desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i)

}

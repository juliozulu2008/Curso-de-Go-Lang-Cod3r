package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou esclusivp nao existe se quiser ter ele ter que fazer assim Var a != var b
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, TV32: %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}

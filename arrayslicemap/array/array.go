package main

import "fmt"

func main() {
	//homogenia (mesmo Tipo) e estatica (fixo) exemplo
	var notas [3]float64 // nome do array tamanha e tipo
	fmt.Println(notas)

	// array é idexado veremos agora como iniciar um array
	notas[0], notas[1], notas[2] = 7.8, 8.9, 5.6
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("Media %.2f\n", media)
}

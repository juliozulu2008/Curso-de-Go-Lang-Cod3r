package main

import "fmt"

func notaParaConceito(n float64) string { // a funçao vai usar uma float64 e ira reornar uma staring
	switch {

	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 3 && n < 0:
		return "E"
	default:
		return "Nenhum valor informado!"
	}
}

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(8.7))
	fmt.Println(notaParaConceito(7.8))
	fmt.Println(notaParaConceito(6.8))
	fmt.Println(notaParaConceito(5.9))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(2.9))
}

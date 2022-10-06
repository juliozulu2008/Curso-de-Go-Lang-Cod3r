package main

import "fmt"

func notaParaConceito(n float64) string { // a funçao vai usar uma float64 e ira reornar uma staring
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
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

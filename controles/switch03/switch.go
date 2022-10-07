package main

import (
	"fmt"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Nao especificado!"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Maoie"))
	fmt.Println(tipo(func() {}))

}

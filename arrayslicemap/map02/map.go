package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Julio Cesar":  3500.50,
		"Kelly Nayara": 2100.50,
		"Mirai Julio":  7890.54,
	}
	funcsESalarios["Joao das Quebradas"] = 1350.0
	delete(funcsESalarios, "Maoie")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}

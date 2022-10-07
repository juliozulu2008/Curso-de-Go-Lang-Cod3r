package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"J": {
			"Julio Cesar": 3590.56,
			"Julio Mirai": 7897.54,
		},
		"K": {
			"Kelly Nayara": 2100.0,
		},
		"F": {
			"Fulano de tal": 9546.5,
		},
	}
	delete(funcsPorLetra, "F")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	fmt.Println(funcsPorLetra)
}

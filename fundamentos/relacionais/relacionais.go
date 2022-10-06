package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String: ", "Banana" == "Banana")
	fmt.Println("Diferente:", 3 != 2)
	fmt.Println("Menor:", 3 < 2)
	fmt.Println("Maior:", 3 > 2)
	fmt.Println("Menor igual:", 3 <= 2)
	fmt.Println("Maior igual:", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}
	p1 := Pessoa{"Julio"}
	p2 := Pessoa{"Julio"}
	fmt.Println("Pessoas:", p1 == p2)
}

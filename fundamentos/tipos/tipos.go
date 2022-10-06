package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//Numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1, i2, i3, i4 := math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64
	fmt.Println("O valor maximo de int é ", i1)
	fmt.Println("O valor maximo de int é ", i2)
	fmt.Println("O valor maximo de int é ", i3)
	fmt.Println("O valor maximo de int é ", i4)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))
	fmt.Println("O tipo de i2 é", reflect.TypeOf(i2))
	fmt.Println("O tipo de i3 é", reflect.TypeOf(i3))
	fmt.Println("O tipo de i4 é", reflect.TypeOf(i4))
	// o valor varia conforme a arquiterura da maquina
	var i5 rune = 'a' // representa um mapeamento da tabela Unicode( int32)
	fmt.Println("O rune é", reflect.TypeOf(i5))
	fmt.Println(i5) // valor referenta a letra A
	// Numeros reais (float32, Float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo de literal de 49.99 é", reflect.TypeOf(49.99))
	// e eu quiser que seja igual tenho que fazer conversao

	//BOLEAN
	bo := true
	fmt.Println("o tipo de Bo é", reflect.TypeOf(bo))
	fmt.Println(bo)
	//STrings
	s1 := "Exemplo de string em GO" // somente com aspas duplas
	fmt.Println(s1)
	fmt.Println("O tamanho da String é ", len(s1))

	//string com multiplas linhas
	s2 := `
	Ola 
	Meu
	Nome
	é 
	Julio`
	fmt.Println(s2)

	// nao tem char eles sao mapeados para int32

}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5} // array
	// slice nao é um array! Slice Define um Pedaço de um array.
	// slice quer dizer pedaço ou fatia em ingles
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:3] //novo slice  mas o mesmo array o a2  neste ele vao do inicio ate a posiçao 2 do index ignorando a 3
	fmt.Println(a2, s3)
	// posso imaginar o slice como : tamanha e um ponteiro para um elemento de um array
	// tambem da pra fazer um slice de um slice coomo no exemplo do s4
	s4 := s2[:1]
	fmt.Println(s2, s4)
}

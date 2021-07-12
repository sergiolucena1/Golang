package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3] int {1, 2, 3} //  array
	s1 := [] int {1, 2, 3} //  slice (pedço do array)
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1),reflect.TypeOf(s1))

	a2 := [5] int {1, 2, 3, 4., 5}

//Slice não é um array! Slice define um pedaço do Array.

	s2 := a2[1:3] // defininado os elementos de indice 1 e 3(sem incluir o 3) do array[2 3]
	fmt.Println(a2, s2)

	s3 := a2[:2] //  vai do indice zero ate o indice 2 ( sem incluir o 2)
	fmt.Println(a2, s3)

// vc pode imaginar o slice como: tamanho e um ponteiro para um elemento de um array

	// slice de um slice

	s4:= s2[:1] // pegando apenas o elemento 0 do slice s2
	fmt.Println(s2, s4)

}
// OBS: todos os slices sao fatias de um só array (a2)

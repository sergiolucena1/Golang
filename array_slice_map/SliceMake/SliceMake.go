package main

import "fmt"

func main(){
	s := make([]int,10) // slice inteiro com 10 elementos
	s[9]= 12 // quero o sline com indice 9 (ultimo elemento) alterando para 12
	fmt.Println(s)

// atribuindo ao slice s inteiro, 10 elementos, mas com um array interno com 20 elementos
	s =  make([]int, 10, 20)

	fmt.Println(s, len(s), cap(s))
// len = tamanho do slice / cap = capacidade

// append = acrescentar no slice (s, esses elementos...)
	s = append(s, 1,2,3,4,5,6,7,8,9,0)
	fmt.Println(s, len(s), cap(s))

// a capacidade do array interno do slice está preenchhida
// so for acrescentar mais algum elemento a capacidade do array dobra de tamanho
	s = append(s,1)
	fmt.Println(s, len(s), cap(s))
	// o array interno passou de 20 elementos para 40 automaticamente

}

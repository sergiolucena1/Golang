package main

import "fmt"

func main (){
	//passando um array(numeros)
	numeros := [...]int {1, 2, 3, 4, 5} // é um array que o compilador vai contar a quantidade de elementos

// i (indice atual) , numero(elemento do array) := range do array
// range numeros (retorna o indice e o elemento do array
	for i, numero := range numeros {
		fmt.Printf("(%d) %d\n", i, numero)
		// me retorna o indice e o numero
	}
// for ignorando o indice do array( _ ) pegando apenas o numero
	for _, num := range numeros{
		fmt.Println(num)
	}
}

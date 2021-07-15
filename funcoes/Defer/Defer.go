package main

import "fmt"

//Defer = adiar a função

func ObterValorAprovado(numero int)int{
	defer fmt.Println("fim!")
	if numero > 5000{
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("valor baixo...")
	return numero
}

func main (){
	fmt.Println(ObterValorAprovado(6000))
	fmt.Println(ObterValorAprovado(3000))
}
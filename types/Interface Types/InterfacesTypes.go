package main

import "fmt"

type curso struct{
	nome string
}

func main (){
	var coisa interface{} // o tipo da variavel é interface
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	//se torna um tipo generico
	type dinamico interface {}

	//string
	var coisa2 dinamico = "opa"
	fmt.Println(coisa2)

	//boolean
	coisa2 = true
	fmt.Println(coisa2)

	//usando struct
	coisa2 = curso {"Golang"}
	fmt.Println(coisa2)
}

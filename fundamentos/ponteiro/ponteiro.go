package main

import "fmt"

func main (){
	i := 1

	var p *int = nil

	p = &i // pegando e endereço da variavel
	*p++    // desreferenciando o ponteiro (pegando o valor da variavel e somando ++)
	i++

	// go nao tem aritmetica de ponteiros
	//p++

	fmt.Println(p, *p , i, &i)
}

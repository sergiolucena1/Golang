package main

import (
	"fmt"
	"strconv"
)

func main(){
	x:= 2.4
	y:= 2
	fmt.Println(x/ float64(y))


	nota:= 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado... ( nao concatena )
	fmt.Println("teste " + string(123)) //  para concatenar com valor inteiro tem q antes transformar em string

	// int para string
	fmt.Println("teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123") // num = primeiro valor  /  _ = erro
	fmt.Println(num - 122) // 123 - 122 = 1

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("verdadeiro")
	}
}

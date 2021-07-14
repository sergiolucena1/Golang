package main

import "fmt"

//armazenar uma func em uma variavel
var soma = func(a,b int)(int){
	return a+b
}

func main (){
	fmt.Println(soma(2,3))

//função local dentro da func main
	sub:= func(a,b int)(int){
		return a-b
	}

	fmt.Println(sub(2,4))
}
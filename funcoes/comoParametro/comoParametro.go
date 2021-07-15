package main

import "fmt"

func multiplicacao(a,b int)int{
	return a*b
}

func exec(funcao func(int,int)int, p1, p2 int)int{
	//funcao func(int,int)int// essa e uma funcao que retorna inteiro
	// func exec (p1,p2 int) int // essa funcao recebe dois parametros int e retorna int

	return funcao(p1,p2)
}

func main (){
	resultado := exec(multiplicacao, 3,4)
	fmt.Println(resultado)
}
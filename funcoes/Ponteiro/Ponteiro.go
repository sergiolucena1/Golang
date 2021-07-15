package main

import "fmt"

func incremento1(n int){
	n++
}

// revisão: um ponteiro é representado por um *

func incremento2 (n *int){
	//revisão: * é usado para desreferenciar, ou seja
	//ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main (){
	n:=1

	incremento1(n) // por valor
	fmt.Println(n)

	// revisão: & usado para obter o endereco da variável
	incremento2(&n) // por referencia
	fmt.Println(n)
}
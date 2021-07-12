package main

import "fmt"

func imprimirResultado (nota float64){
	if nota >= 7 {
		fmt.Println("O aluno foi aprovado com a nota :", nota )
	}else {
		fmt.Println("reprovado com a nota :", nota )
	}
}

func main (){
	imprimirResultado(7.3)
	imprimirResultado(5.3)
}
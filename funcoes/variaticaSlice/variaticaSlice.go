package main

import "fmt"

func ImprimirAprovados(aprovados...string){
	fmt.Println("lista de Aprovados")
	for i, aprovados := range aprovados{
		//i = indice do slice / aprovados = elemento
		fmt.Printf("%d %s\n",i+1 ,aprovados)
	}
}

func main(){
	aprovados := []string{"Maria", "Pedro", "Rafael", "Raissa"}//slice
	ImprimirAprovados(aprovados...)//usa os parametros para a funcao
}
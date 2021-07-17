package main

import "fmt"

func main(){
	ch := make(chan int, 1)// criando um canal que dentro dele, será enviado valores inteiros

	ch <- 1 //Envindo dados para um canal (escrita)
	<-ch // recebendo dados do canal (leitura

	ch <- 2
	fmt.Println(<-ch) // aqui eu leio os dados do canal
}

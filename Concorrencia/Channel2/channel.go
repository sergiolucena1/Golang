package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre as goroutines(sincroniza os dados)
// channel é um tipo da linguagem

func doisTresQuatroVezes(base int, c chan int){
	time.Sleep(time.Second)
	c<- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c<- 3 * base

	time.Sleep(time.Second)
	c<- 4 * base
}

func main (){
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // recebendo os dados do canal (leitura dos dados)
	fmt.Println(a,b)

	fmt.Println(<-c)
}
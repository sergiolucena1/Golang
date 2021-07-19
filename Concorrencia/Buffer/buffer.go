package main

import (
	"fmt"
	"time"
)

func rotina (ch chan int){
	ch<-1
	ch<-2
	ch<-3
	ch<-4 // buffer fica lotado aqui (até aqui ele executa)
	fmt.Println("Executou!") // nao vai executar essa linha pois nao cabe no buffer
	ch<-5
	ch<-6
}

func main(){
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
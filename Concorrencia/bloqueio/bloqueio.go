package main

import (
	"fmt"
	"time"
)

func rotina (c chan int){
	time.Sleep(time.Second) // esperar 1 segundo
	c<- 1 // operção bloqueante (pq nao tem buffer)
	fmt.Println("só depois que foi lido")
}

func main (){
	c:= make(chan int)// canal sem buffer

	go rotina (c)

	fmt.Println(<-c)// operação bloqueante
	fmt.Println("foi lido")
	fmt.Println(<-c) // deadlock (pois nao tem mais buffer para executar as goroutines e os canais)
	fmt.Println("Fim")
}
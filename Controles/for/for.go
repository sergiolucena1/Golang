package main

import (
	"fmt"
	"time"
)

func main (){
	i:= 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for i := 0 ; i <= 20; i+=2{ // for inline até 20 pulando de 2 em 2

		fmt.Printf("%d ", i)
	}
	fmt.Println("\n Misturando...")
	for i :=1 ; i <= 10 ; i ++{
		if i%2 == 0  {
			fmt.Print("Par ")
		}else{
			fmt.Print("Ímpar ")
		}
	}
	for{
		//laço Infinito
		fmt.Println("Para sempre ")
		time.Sleep(time.Second * 5) // no espaço de tempo de 5 segundo
	}
}

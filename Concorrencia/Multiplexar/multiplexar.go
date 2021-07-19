package main

import (
	"fmt"
	"github.com/html"
)

// func encaminhar: recebe um dado de um canal de origem e encaminha para um outro canal de destino

func encaminhar(origem <-chan string, destino chan string){
	for{
		destino <- <- origem //sempre que chegar dados no canal origem eu vou passar pro destino
	}
}

//multiplexar : misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <- chan string{
	c := make (chan string)
	go encaminhar(entrada1, c) // pegar a entrada1 e encaminha pro canal c
	go encaminhar(entrada2, c)

	return c
}

func main (){
	c := juntar( // juncao dos dois canais
		html.Titulo("https://cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.amazon.com", "https://www.youtube.com"),
		)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)

}

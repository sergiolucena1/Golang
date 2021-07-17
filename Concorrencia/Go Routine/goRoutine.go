package main

import (
	"fmt"
	"time"
)

//passar o nome da pessoa, o texto q ela vai falar e a quantidade de vezes q esse texto vai ser impreso
func fale(pessoa, texto string, qtde int){
	for i:= 0; i<qtde; i++ {
		 time.Sleep(time.Second)// vai passar 1 segundo parado e depois...
		 fmt.Printf("%s: %s (interação %d)\n",pessoa,texto,i+1)//interação informa a quantidade de vezes que a pessoa fala
	}
}

func main (){
//processo de vc chamar funçoes sem o uso da concorrência (forma serial sem ser independente)
	//fale("maria","Pq vc não fala comigo?",3)
	//fale("João.","Só posso falar depois de vc",1)

// usar "go" antes de chamar o método criar uma go routine( de forma independente)
//	go fale("Maria","Ei...",500)
//	go fale("João", "Opa...",500)

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabéns", 5)

}
//Obs: Muito parecido com a criação de Thread(ordem de execução)
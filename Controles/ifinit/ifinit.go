package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAletorio () int {
	s := rand.NewSource(time.Now().UnixNano()) // rand ( numero aleatorio que tem como fonte(source)   --- time.now(data atual) -- UnixNano = nano segundo
	r := rand.New(s)
	return r.Intn(10) // gera um numero aleatorio ate 10
}

func main (){
	if i := numeroAletorio(); i > 5 { // tbm é suportado no switch
		fmt.Println("Ganhou")
	}else {
		fmt.Println("perdeu")
	}
}
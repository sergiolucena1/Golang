package main

import (
	"fmt"
	"time"
)
// função para me saldar dependendo da hora do dia
func main () {
	t :=  time.Now() // hora atual
	switch {// switch true(procura alguma case que seja verdadeiro
	case t.Hour() < 12:
		fmt.Println("Bom dia !")
	case t.Hour() <18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa Noite ")
	}
}


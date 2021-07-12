package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string{   // interface{} é uma variavel genérica pode receber qualquer tipo
	switch i.(type) {
	case int:
		return "inteiro"
	case float64, float32:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}
}

func main () {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(2))
	fmt.Println(tipo("Sérgio"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))




}
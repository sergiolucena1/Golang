package main

import "fmt"

type esportivo interface{
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//pode adicionar mais métodos
}

type bmw7 struct {}

// bmw7 e receiver dos dois types

func (b bmw7) ligarTurbo(){
	fmt.Println("turbo...")
}
func (b bmw7) fazerBaliza(){
	fmt.Println("Baliza")
}

func main (){
//executando o codigo a partir de uma interface que é a composicao de outras
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
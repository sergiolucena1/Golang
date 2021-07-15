package main

import "fmt"

type carro struct {
	nome string
	velocidadeAtual int
}

type Ferrari struct {
	carro //campo anonimo (sem identificador)
	turboLigado bool // atributo especifico
}

func main(){
	f := Ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado )
	fmt.Println(f)
}
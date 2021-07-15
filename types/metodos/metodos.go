package main

import (
	"fmt"
	"strings"
)

type pessoa struct{
	nome string
	sobrenome string
}

func(p pessoa) getNomeCompleto()string{
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCempleto string){
		//* ponteiro pra toranar a variavel mais abrangente

	partes := strings.Split(nomeCempleto," ")
//dividindo o nome completo (.Split(nomecompleto) e armazenando na var partes

	p.nome = partes[0]//o index zero da var partes
	p.sobrenome = partes[1] // indice 1 da var partes
}

func main (){
	//funcao get
	p1 := pessoa{"Sérgio", "Lucena"}
	fmt.Println(p1.getNomeCompleto())

	//funcao set
	p1.setNomeCompleto("Roberto Torres")
	fmt.Println(p1.getNomeCompleto())
}
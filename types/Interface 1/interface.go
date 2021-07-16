package main

import "fmt"

// toda interface(imprimivel) tem q ter o metodo toString retornando uma string

type imprimivel interface{
	toString() string
}

type pessoa struct {
	nome string
	sobrenome string
}

type produto struct {
	nome string
	preco float64
}

//Interfaces são implementadas implicitamente(oculta)
func (p pessoa) toString() string{
	return p.nome + " " + p.sobrenome
}

func (p produto) toString()string{
	return fmt.Sprintf("%s - R$%.2f", p.nome, p.preco)
}

func imprimir (x imprimivel){  // recebe como parametro qualquer coisa que seja imprimivel
	fmt.Println(x.toString())
}

func main(){
	var coisa imprimivel = pessoa{"Sergio","Lucena"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	//polimorfismos da mesma variavel imprimivel
	coisa = produto{"Calça Jeans", 88.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	//ou passar diretamento com a função imprimir
	produto2:= produto{"Bermuda",44.30}
	imprimir(produto2)
}
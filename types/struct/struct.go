package main

import "fmt"

//struct = estrutura (agrupamento de dados)
//criando a estrutura
type produto struct {
	nome string
	preco float64
	desconto float64
}

// Método: função com receiver(receptor) ou seja func que recebe o type
//func (receiver) (parametros)(tipo de retorno)
//p = recebe o type(p.preco: acessa o atributo da struct)
func (p produto) precoComDesconto()float64{
	return p.preco * (1- p.desconto)
}

func main (){
//criando uma var e atribuindo um produto ao type
	var produto1 produto
	produto1 = produto{
		nome: "Lápis",
		preco: 1.88,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

//Outra forma mais reduzida de atribuir um produto ao tipo
	produto2:= produto{"notebook",2300.50, 0.08}
	fmt.Println(produto2.precoComDesconto())
}
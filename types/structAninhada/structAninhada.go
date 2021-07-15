package main

//struct dentro de outro

import "fmt"

type item struct{
	produtoID int
	qtd int
	preco float64
}

type pedido struct{
	userID int
	itens []item //slice de item
}
//receiver pedido
func (p pedido) valorTotal()float64{
	total:= 0.0
	for _, item := range p.itens{
		total += item.preco * float64(item.qtd) // converter para float64
	}	//o total somou e atribuiu(+=) o preço * a quantidade de itens
	return total
}

func main(){
	pedido := pedido{
		userID: 1,
		itens:[]item{
			item{1, 2,22.50},
			item{2,1,44.30},
			item{11,100,4.55},
		},

	}
	fmt.Printf("Valor total do pedido é %.2f", pedido.valorTotal())
}
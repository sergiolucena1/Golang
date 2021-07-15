package main

import "fmt"

func fatorial (n uint)uint{ //forcando que as entradas sejam >0 uint
	switch  {
	case n== 0:
		return 1  // nao vai ter erro
	default:

		return n * fatorial(n-1) // definicao do fatorial
	}
}
func main () {
	resultado := fatorial(5) //chamada valida
	fmt.Println(resultado)
}
	//uma solução melhor seria ...

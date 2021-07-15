package main

import "fmt"

func fatorial (n int)(int, error){
	switch  {
	case n <0: //caso for um numero negativo
		return  -1, fmt.Errorf("Numero inválido: %d", n) // da erro
	case n== 0:
		return 1 , nil // nao vai ter erro
	default:
		fatorialAnterior,_:=fatorial(n -1) //chamada recursiva
		return n * fatorialAnterior, nil // definicao do fatorial
	}
}
func main (){
	resultado, _:= fatorial(5) //chamada valida
	fmt.Println(resultado)

	_,err := fatorial(-4) // chamada inválida
	if err != nil{
		fmt.Println(err)
	}
	//uma solução melhor seria ...
}
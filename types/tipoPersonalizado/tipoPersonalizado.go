package main

import "fmt"

type nota float64

func (n nota ) entre(notaInicio, notaFinal float64)bool {
	return float64(n) >= notaInicio && float64(n)<= notaFinal //  se a nota esta entre o inicio e o fim

}

func notaParaConceito (n nota )string	{
	if n.entre(9.0,10.0){
		return "A"
	}else if n.entre(7.0,8.99){
		return "B"
	}else if n.entre(5.0, 7.99) {
		return "C"
	}else if n.entre(3.0,4.99) {
		return "D"
	}else {
		return "E"
	}
}

func main (){
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(6.5))
	fmt.Println(notaParaConceito(2.1))
}
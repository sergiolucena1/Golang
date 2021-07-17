package main

import "fmt"
// acessando do arquivo reta.go, no mesmo pacote
func main(){
	p1:= Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1,p2))
}

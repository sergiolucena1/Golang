package main

import "fmt"

func f1 (){
	fmt.Println("F1")
}
func f2 (p1 string, p2 string){ // parametros
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3()string{
	return "F3"
}

func f4(p1, p2 string)string{
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

//funcao que n recebe parametro e que retorna dois valores
func f5 () (string, string){
	return "retorno1","retorno 2"
}

//chamada de cada uma das funcoes
func main (){
	f1()
	f2("Param1", "Param2")

	//retorno f3 e f4
	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	//retorno f5 retorno multiplo
	r51, r52 := f5()
	fmt.Println("F5", r51, r52)
}
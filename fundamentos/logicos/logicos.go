package main

import "fmt"

func compras(trab1, trab2 bool)(bool, bool, bool ){
	comprarTv50 := trab1 && trab2 // (and)
	comprarTv32 := trab1 != trab2 //  ou exclusivo
	comprarSorvete := trab1 || trab2 // ou

	return comprarTv50, comprarTv32, comprarSorvete

}
func main (){							// alterar aqui outras opções logicas
	tv50 ,  tv32 , sorvete := compras(true , true ) //  passando q os dois trabalhos deram certo (true e true )
	fmt.Printf("Tv50: %t, Tv 32: %t , Sorvete: %t ,  Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
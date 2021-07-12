package main

import "fmt"

func main() {
	//homogenia (mesmo tipo) e estatica(fixo) vc decide na inicialização do array seu tamanho
	var notas [3] float64  // 3 é o tamanho do array , float é o tipo
	fmt.Println(notas)
	// os elementos do array começam no indice 0
	notas[0], notas[1], notas[2] = 7.8, 4.3 , 9.1 // agregando valor aos indices
	fmt.Println(notas)

	total := 0.0
	for i := 0 ; i< len(notas) ; i ++ { // len= mostra o tamanho do array
		total += notas[i] // total é a soma de todos os elementos do array
		fmt.Println(total)
	}
	// a variavel total é do tipo float64, o len é do tipo int( entao temos que converter para dividir
	// a media do array
	media := total / float64(len(notas)) // convertendo o tamanho do array
	fmt.Printf("Média %.2f\n", media) // mostrar float apenas com duas casas decimais (2f)
}

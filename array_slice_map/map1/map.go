package main

import "fmt"

func main (){
	//var aprovados map[int] string
	// mapas devem ser inicializados

						//[chave]valor
	aprovados := make(map[int]string) // map com chave inteira e valor string

	aprovados [111222333444555] = "maria"
	aprovados [234234455634545] = "João"
	aprovados [887899965578788] = "raissa"
	fmt.Println(aprovados)

//for com map
	for cpf , nome := range aprovados{
		fmt.Printf("%s (CPF: %d\n", nome, cpf) // %s = string  %d = inteiro

	}
	//ler um valor do map a partir da chave
	fmt.Println(aprovados[234234455634545])
	//deletar um valor do map usando a chave
	delete(aprovados, 234234455634545 )
	fmt.Println(aprovados)
}

package main

import "fmt"

func main(){
	// o valor do primeiro map é outro map que retorar um float64
	funcionariosPorLetra:= map[string]map[string]float64{
		"G":{
			"Gabriela Silva": 23444.88,
			"Guga Pereira": 3452.33,
		},
		"J":{
			"José João":11234.65,
		},
		"P":{
			"Pedro Junior": 1200.0,
		},
	}
	delete(funcionariosPorLetra, "P") // deletei todos os funcionario com a letra "P"

	for letra , funcs := range funcionariosPorLetra{
		fmt.Println(letra, funcs)
	}
}

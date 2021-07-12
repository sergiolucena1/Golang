package main

import "fmt"

func main (){
	//iniciando o map ja com os valores de forma literal
	funcsEsalarios := map[string]float64{
		"Sérgio Lucena": 23444.23,
		"Raissa Mahon": 15250.20,
		"Pedro Henrique": 22321.44,
	}
	//adicionando valor depois
	funcsEsalarios["Diva Galdencia"] = 5430.0
	//deletando um elemento
	delete(funcsEsalarios, "Inexistente")

	for nome, salario := range funcsEsalarios{
		fmt.Println(nome, salario)
	}
}

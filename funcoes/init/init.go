package main

import "fmt"

//executa antes da função main
func init(){
	fmt.Println("Inicializando...")
}

func main (){
	fmt.Println("Main...")
}

//dentro de cada arquivo go vc pode ter uma funcao init que é executada antes do main
//também pode ser compilada e usada dentro do arquivo pelo terminal

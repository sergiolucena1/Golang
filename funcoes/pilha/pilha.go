package main

import "runtime/debug"

//pilha de funções termina e começa no main
func f3(){
	// imprimir a pilha de execusao de um programa
	debug.PrintStack()
}

func f2(){
	f3()
}
func f1() {
	f2()
}

func main(){
	f1()
}
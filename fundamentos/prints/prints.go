package main

import "fmt"

func main () {
	fmt.Print("primeira linha ") //  fica na mesma linha
	fmt.Print("primeira linha ")

	fmt.Println("ainda na primeira linha ") //  println (quebra linha) faz q a prixima chamada va pra proxima linha
	fmt.Println("Segunda linha  ")

	x:= 3.141516
	//concatenando
	//fmt.Println("O valor de x é" + x)   ( vai da erro)
	xs := fmt.Sprint(x)   // retorna em formato de string
	fmt.Println("O valor de x é :"+ xs)

	// print formatado
	fmt.Printf("O valor de x é %.2f", x) // .2f (apenas duas casas decimais apos o ponto) sendo do tipo float(f)

	a:=1
	b:= 1.9999
	c:= false
	d:= "sergim"
	fmt.Printf("\n%d %f %.1f %t %s", a , b,b ,c ,d )
						// \n 9quebra linha, %d(inteiro), %f(float), %t(boolean) , %s(string)
	// ou usar %v ( que serve pra varios tipos diferentes
}
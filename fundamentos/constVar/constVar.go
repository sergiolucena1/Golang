package main

import (
	"fmt"
	"math"
)

func main() {
	const PI  float64 = 3.1415 //  const, o nome da comstante(pi) e o tipo da constante (float64)
	var raio = 3.2 // tipo (float 64) inferido pelo compilador

	// forma reduzida de criar uma variavel
	area := PI * math.Pow(raio,2) // :=  significa declarar e atribuir
	// sempre tem q chamar a variavel
	fmt.Println("A area da circunferencia é:", area )

	const(
		a =1
		b =2
	)
	var(
		c = 3
		d = 4
	)
	fmt.Println(a,b,c,d)

	// declarando variaveis em uma linha
	var e, f  bool = true ,  false
	fmt.Println(e, f )

	//ou
	g, h, i := 2 , false , "sérgio" // g= 2 , h = flase , i = "sergio"
	fmt.Println(g,h,i)
}

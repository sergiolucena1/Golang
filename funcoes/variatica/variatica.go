package main

import "fmt"

//...(parametro variatico)
func media (numeros ...float64)float64{
	total := 0.0
	for _, num := range numeros{
		total += num
	}
	return total / float64(len(numeros))//media
}

func main (){
	fmt.Printf("Média: %2.f", media(7.7,9.2,4.3,8.7))
}									// posso passar quantos parametros quiser
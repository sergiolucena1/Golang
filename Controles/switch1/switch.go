package main

import "fmt"

func NotaParaConceito(n float64) string{
	var nota = int(n)
	switch nota {
	case 10:     // se a nota for 10 fallthrough(nao faça nada vá para o proximo)
		fallthrough
	case 9:  // caso a nota for 9...
		return "A"
	case 8 , 7 :
		return "B"
	case 6 , 5 :
		return "c"
	case 4 , 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"

	}
}
func main (){
	fmt.Println(NotaParaConceito(10))
	fmt.Println(NotaParaConceito(8))
	fmt.Println(NotaParaConceito(6))
	fmt.Println(NotaParaConceito(3))
	fmt.Println(NotaParaConceito(33))


}
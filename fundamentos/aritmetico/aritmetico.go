package main

import (
	"fmt"
	"math"
)

func main () {
	a := 3
	b := 2
	c := 3.0
	d:= 2.0

	//operações usando Math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b))) //  tive que converter int para float
	fmt.Println("Menor=>", math.Min(c , d))
	fmt.Println("exponencial =>", math.Pow(c , d))
}
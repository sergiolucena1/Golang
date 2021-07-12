package main

import (
	"fmt"
	"math"
	"reflect"
)

func main () {
	// numeros inteiros
	fmt.Println(1, 2, 10000)
	fmt.Println("literal inteiro é :", reflect.TypeOf(3200000))

	// sem sinal (so valores positivos)... uint8 uint16 uint32 uint64
	var b  byte = 255
	fmt.Println("o byte é ", reflect.TypeOf(b))
	i1 := math.MaxInt64
	fmt.Println("o valor maximo do int é:" , i1)
	fmt.Println("o tipo i1 é:", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa  um mapeamento da tabela unuicode (int32)
	fmt.Println("o rune é :",reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float 32 , float 64)
	var x float32 = 49.99
	fmt.Println("o tipo de x é:", reflect.TypeOf(x))
	fmt.Println("o tipo literal 49.99 é : ", reflect.TypeOf(49.99))

	// boolean
	bo:= true
	fmt.Println( "o tipo de bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "olá, meu nome é sergio"
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho da string é :", len(s1))

}

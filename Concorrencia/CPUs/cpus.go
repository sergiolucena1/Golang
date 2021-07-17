package main

import (
	"fmt"
	"runtime"
)

// informa quantos processadores físicos sua máquina tem

func main(){
	fmt.Println(runtime.NumCPU())
}

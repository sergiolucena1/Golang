package main

//convertendo struct para Json e json para struct

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
	Tags []string `json:"tags"`
}
func main(){
	// struct para json
	p1:= produto{1,"notebook", 1570.00,[]string{"Promoção","Eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json)) //  ja retorna em json

	//json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":5.65,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
	fmt.Println(p2.Preco)

}
package main

import "math"

// Inicializando com letra maiuscula é Público(visível dentro e fora do pacote)
// inicializando com letra minuscula é Privado(visível apenas dentro do pacote)

// Por exemplo...
// Ponto - gerará algo público
// ponto ou _Ponto - gerará algo privado


// Quando a função for pública deve ter um comentario informando mais sobre ela...

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct { //struct publica
	x float64
	y float64
}

// func cateto é visivel apenas dentro do pacote(comeca com letra minuscula)Privada
func catetos(a, b Ponto) (cx, cy float64){ //cateto x , cateto y
	cx = math.Abs(b.x - a.x) // ABS = valor absoluto
	cy =  math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto)float64{ // metodo publico
	cx, cy := catetos(a,b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy,2))
}



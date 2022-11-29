package main

import (
	"fmt"
)

func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}

func main() {
	fmt.Println(Somando(1))
	fmt.Println(Somando(1, 1))
	fmt.Println(Somando(1, 1, 1))
	fmt.Println(Somando(1, 1, 2, 4))
}

/*Note o uso das reticências na declaração do parâmetro número: numeros ...int.
Portanto, podemos criar uma função sem parâmetro, com um, dois, três,
ou uma quantidade indeterminada de parâmetros com Go.*/

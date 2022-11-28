package main

import (
	"fmt"
	"reflect"
)

func main() {
	exibeNomesArray()
	fmt.Println("-------------------------")
	exibeNomes()
}

// array de string
// Em GO não trabalhamos com Arrays pq temos algumas desvantagens, uma delas é ter que declarar de forma fixa o tamanho do array
// Então no GO nós trabalhamos com outra estrutura chamada slice, ele é como se fosse uma abstração do array
func exibeNomesArray() {
	var nomes [4]string
	nomes[0] = "Douglas"
	nomes[1] = "Daniel"
	nomes[2] = "Bernardo"

	fmt.Println(nomes)
	fmt.Println("O meu array tem", len(nomes), "itens")
	fmt.Println("O meu array tem capacidade para", cap(nomes), "itens")
	fmt.Println(reflect.TypeOf(nomes))
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println(nomes)
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida") //add uma nova posiçao no array slice
	fmt.Println(nomes)
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens") // Toda vez que adiciono uma nova posição no slice a capacidade dele duplica, por isso apareceu 6 itens

	fmt.Println(reflect.TypeOf(nomes))
}

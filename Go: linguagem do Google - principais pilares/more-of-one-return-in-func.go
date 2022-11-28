package main

import "fmt"

func main() {
	nome := devolveNome()
	fmt.Println(nome)

	nome2, idade := devolveNomeEIdade()
	fmt.Println(nome2, "e tenho ", idade, "anos")

	nome3, sobrenome, idade2 := devolveNomeSobrenomeEIdade()
	fmt.Println(nome3, sobrenome, idade2)

	// Quando eu não for utilizar todas as variaveis da funçao posso declarar a posição da mesma com um _
	nome4, sobrenome1, _ := devolveNomeSobrenomeEIdade()
	fmt.Println(nome4, sobrenome1)
}

// função retorna um valor
func devolveNome() string {
	nome := "Ingrid"
	return nome
}

// função retorna 2 valores
func devolveNomeEIdade() (string, int) {
	nome := "Ingrid"
	idade := 24
	return nome, idade
}

// função retorna 3 valores mas poderia retornar mais
func devolveNomeSobrenomeEIdade() (string, string, int) {
	nome := "Ingrid"
	sobrenome := "Paulino"
	idade := 24
	return nome, sobrenome, idade
}

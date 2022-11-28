package main

import (
	"fmt"
)

func main() {
	exibeMenuDeEscolha()
	comando := leComandoDigitado()

	if comando == 1 { //No GO condição sempre tem que retornar booleano (truer or false)
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conheço este comando")
	}

	fmt.Println("--------")

	switch comando { //No GO não tem break
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço este comando")
	}
}

func exibeMenuDeEscolha() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComandoDigitado() int { //int é o tipo do retorno da func
	var comandoLido int
	// Capitura o que o usuario escreve
	fmt.Scanf("%d", &comandoLido)
	//ou
	//fmt.Scan(&comando) //nao preciso passar o modificador
	// %d sig que espero um inteiro
	//& sig que oq o usuario digitar eu quero salvar na variavel comando
	fmt.Println("O endereço da minha variavel comando é", &comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	fmt.Println("")

	return comandoLido
}

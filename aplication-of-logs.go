package main

import "fmt"

func main() {
	nome := "Ingrid"
	versao := 1.1
	fmt.Printf("Olá", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	//Capitura o que o usuario escreve
	fmt.Scanf("%d", &comando)
	//ou
	//fmt.Scan(&comando) //nao preciso passar o modificador
	// %d sig que espero um inteiro
	//& sig que oq o usuario digitar eu quero salvar na variavel comando
	fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("O comando escolhido foi:", comando)

	if comando == 1 { //No go condição sempre tem que retornar booleano (trur or false)
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conheço este comando")
	}
}

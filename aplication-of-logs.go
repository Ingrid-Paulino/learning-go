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
	fmt.Scan(&comando) //nao preciso passar o modificador
	// %d sig que espero um inteiro
	//& sig que oq o usuario digitar eu quero salvar na variavel comando
	fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("O comando escolhido foi:", comando)
}

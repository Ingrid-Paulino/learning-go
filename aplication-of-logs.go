/*
Essa aplicaçao monitora sites para ver se os mesmos estão online ou não
*/
package main

import (
	"fmt"
	"net/http" //pacote que faz requisição web
	//É importante saber que temos vários subdiretórios dentro do "net". Se quiséssemos fazer um envio de email poderíamos usar o "net/smtp".
	"os" //pacote que comunica com o sistema operacional
)

func main() {
	exibeIntroducao()
	exibeMenu()

	//var comando int
	//comando = leComando()
	//ou
	comando := leComando()

	//if comando == 1 { //No GO condição sempre tem que retornar booleano (truer or false)
	//	fmt.Println("Monitorando...")
	//} else if comando == 2 {
	//	fmt.Println("Exibindo Logs...")
	//} else if comando == 0 {
	//	fmt.Println("Saindo do programa...")
	//} else {
	//	fmt.Println("Não conheço este comando")
	//}

	// ou

	switch comando { //No GO não tem break
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0) //retorna status 0, o codigo 0 informo para o sistema operacional que ele saiu com sucesso
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1) // -1 indico que teve algum problema para o sistema operacional na hora de sair o programa
	}
}

func exibeIntroducao() {
	nome := "Ingrid"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int { //int é o tipo do retorno da func
	var comandoLido int
	// Capitura o que o usuario escreve
	fmt.Scanf("%d", &comandoLido)
	//ou
	//fmt.Scan(&comando) //nao preciso passar o modificador
	// %d sig que espero um inteiro
	//& sig que oq o usuario digitar eu quero salvar na variavel comando
	fmt.Println("O endereço da minha variavel comando é", &comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	//resp, err := http.Get(site) //acessa o site //get retorna uma resposta=resp ou um erro=err
	resp, _ := http.Get(site)
	fmt.Println(resp)
}

// Introdução a GO
package main /*Informa que o programa vai ser o principal pacote da aplição e que seu codigo precisa
começar por ele*/

import "fmt"

func main() { // função principal
	fmt.Println("Hello, World!") //Em go todas as funçoes de pacotes externos comecam com letra maiuscula
}

/*
No terminal rode o go:
	- go build nomeDoarquivo.go
		- Um arquivo executável é gerado
			- ./nomeDoArquivo
Para não ter que sempre fazer o processo acima a cada alteracao no arquivo
posso rodar o comando --> go run nomeArquivo.go (compila e executa)
*/

// Não precisa de ; em go
// O nome da função principal tem que ser sempre main
// pode ter mais de uma função principal

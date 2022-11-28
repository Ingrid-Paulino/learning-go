package main

import (
	"fmt"
)

func main() {
	//printNamesForma1()
	fmt.Println("")
	printNamesForma2()
	fmt.Println("")
	printNamesForma3()
}

// Quando passo o for sem nada, sig. que esse for vai ficar em loop eternamente
// nessa aplicação para eu sair do programa preciso colocar a opção 0 para parar o loop
// No  Go nao tem while
//func printNamesForma1() {
//exibeIntroducao()
//for {
//	exibeMenu()
//	comando := leComando()
//
//	switch comando { //No GO não tem break
//	case 1:
//		fmt.Println("Monitorando...")
//	case 2:
//		fmt.Println("Exibindo Logs...")
//	case 0:
//		fmt.Println("Saindo do programa...")
//		os.Exit(0) //retorna status 0, o codigo 0 informo para o sistema operacional que ele saiu com sucesso
//	default:
//		fmt.Println("Não conheço este comando")
//		os.Exit(-1) // -1 indico que teve algum problema para o sistema operacional na hora de sair o programa
//	}
//}
//}

func printNamesForma2() {
	nomes := []string{"Ingrid", "Jacke", "Gustavo"}
	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i]) //printa 1 de cada vez
	}
}

func printNamesForma3() {
	nomes := []string{"Ingrid", "Jacke", "Gustavo"}

	for i, nome := range nomes { //range: operaçao de interação do GO, da acesso a cada item do array
		// range pode me retornar duas coisas: 1- qual a posição que ele esta interando e 2- o que esta na posição
		fmt.Println("Meu nome esta na posição", i, ":", nome)
	}
}

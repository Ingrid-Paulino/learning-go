package main

import "fmt"

// Estrutura de conta corrente
type ContaCorrente struct { //struct sinaliza que é uma estrutura
	// cada atributa é uma variavel
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// Forma 1
	fmt.Println(ContaCorrente{}) // Quando não passo nenhum valor para a struct ela me retorna "" vazias e 0 na posiçao de cada atributo no obj
	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	//ou
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	//ou
	//nao preciso passar todos os atributos
	contaDaIngrid := ContaCorrente{titular: "Ingrid", saldo: 125.5}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDaIngrid)
	fmt.Println("--------------------")

	// Forma 2
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(contaDaCris) // printa com o &
	fmt.Println(*contaDaCris)
}

//essa forma de escrita ele compara o conteudo dessa comparação

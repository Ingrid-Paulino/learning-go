package main

import "fmt"

type ContaCorrente2 struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	//essa forma de escrita ele compara o conteudo dessa comparação
	contaDoGuilherme := ContaCorrente2{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	contaDoGuilherme2 := ContaCorrente2{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	fmt.Println(contaDoGuilherme == contaDoGuilherme2) // retorna true

	contaDaBruna := ContaCorrente2{"Bruna", 222, 111222, 200}
	contaDaBruna2 := ContaCorrente2{"Bruna", 222, 111222, 200}
	fmt.Println(contaDaBruna == contaDaBruna2) // retorna true

	fmt.Println("---------------")

	//essa forma de escrita ele compara o endereço da memoria e o conteudo
	var contaDaCris *ContaCorrente2
	contaDaCris = new(ContaCorrente2)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	var contaDaCris2 *ContaCorrente2
	contaDaCris2 = new(ContaCorrente2)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500

	fmt.Println(contaDaCris)
	fmt.Println(contaDaCris2)
	//mostra o local na memoria
	fmt.Println(&contaDaCris)
	fmt.Println(&contaDaCris2)

	//essa forma de escrita ele compara o endereço de memoria
	fmt.Println(contaDaCris == contaDaCris2) // retorna false
	//essa forma de escrita ele compara o conteudo
	fmt.Println(*contaDaCris == *contaDaCris2) // retorna true

}

package main

import (
	"fmt"
	"learning-go/Go: Orientacao a Objetos/gerencia contas corrente/contas"
)

// conta verificarConta: Verifica se no parametro que eu passo tem a função sacar
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

//Tipos polimórficos em Go
/*
Uma interface é a definição de um conjunto de métodos comuns a um ou mais tipos. É o que permite a criação de tipos polimórficos em Go.
Java possui um conceito muito parecido, também chamado de interface. A grande diferença é que, em Go, um tipo implementa uma interface implicitamente.
*/
//Podemos reaproveitar métodos ou funções entre tipos utilizando interface.
// Nosso projeto é um exemplo disso. Reutilizamos o método Sacar para pagar os boletos com ambas as contas.
type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60) // passamos o & para ele saber qual o endereço que estou buscando

	fmt.Println(contaDoDenis.ObterSaldo())
}

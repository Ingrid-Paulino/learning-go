package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// (c *ContaCorrente): é como se fosse o "this", c vai seu o objeto da pessoa que estiver tentando sacar o dinheiro
// c é quem esta chamando a função sacar
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(400)) // c na função sacar é contaDaSilva
	fmt.Println(contaDaSilvia.saldo)

}

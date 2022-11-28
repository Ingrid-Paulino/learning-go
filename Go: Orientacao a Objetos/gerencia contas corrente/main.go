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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(400)) // c na função sacar é contaDaSilva
	fmt.Println(contaDaSilvia.saldo)

	//Forma 1
	fmt.Println(contaDaSilvia.Depositar(2000))
	fmt.Println(contaDaSilvia.saldo)

	//Forma 2
	fmt.Println(contaDaSilvia.saldo)
	status, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(status, valor)

	//Transferencia
	contaDaSilvia2 := ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo2 := ContaCorrente{titular: "Gustavo", saldo: 100}

	status2 := contaDaSilvia2.Transferir(200, &contaDoGustavo2) //&contaDoGustavo: & indica que é para essa conta que o dinheiro sera transferido

	fmt.Println(status2)
	fmt.Println(contaDaSilvia2)
	fmt.Println(contaDoGustavo2)
}

package main

import (
	"fmt"
	"learning-go/Go: Orientacao a Objetos/gerencia contas corrente/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{}
	//contaDaSilvia.Titular = "Silvia"
	//saldo nao vai funcionar pois ele é privado no outro arquivo, para fazer a leitura dele, eu teria que chamar a funcao ObterSaldo()
	//contaDaSilvia.saldo = 500
	fmt.Println(contaDaSilvia)

	//fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(400)) // c na função sacar é contaDaSilva
	//fmt.Println(contaDaSilvia.saldo)

	//Forma 1
	fmt.Println(contaDaSilvia.Depositar(2000))
	//fmt.Println(contaDaSilvia.saldo)

	//Forma 2
	//fmt.Println(contaDaSilvia.saldo)
	status, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(status, valor)

	//Transferencia
	//contaDaSilvia2 := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	//contaDoGustavo2 := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	//status2 := contaDaSilvia2.Transferir(200, &contaDoGustavo2) //&contaDoGustavo: & indica que é para essa conta que o dinheiro sera transferido

	//fmt.Println(status2)
	//fmt.Println(contaDaSilvia2)
	//fmt.Println(contaDoGustavo2)
}

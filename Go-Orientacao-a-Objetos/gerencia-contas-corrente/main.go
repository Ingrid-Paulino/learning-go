package main

import (
	"fmt"
	"learning-go/Go: Orientacao a Objetos/gerencia-contas-corrente/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{}
	//contaDaSilvia.Titular = "Silvia"
	//saldo nao vai funcionar pois ele é privado no outro arquivo, para fazer a leitura dele, eu teria que chamar a funcao ObterSaldo()
	//contaDaSilvia.saldo = 500
	fmt.Println(contaDaSilvia)

	//fmt.Println(contaDaSilvia.saldo)

	//Forma 1
	fmt.Println(contaDaSilvia.Depositar(2000))
	//fmt.Println(contaDaSilvia.saldo)

	//Forma 2
	//fmt.Println(contaDaSilvia.saldo)
	//status, valor := contaDaSilvia.Depositar(2000)
	//fmt.Println(status, valor)

	fmt.Println(contaDaSilvia.Sacar(400)) // c na função sacar é contaDaSilva
	//fmt.Println(contaDaSilvia.saldo)

}

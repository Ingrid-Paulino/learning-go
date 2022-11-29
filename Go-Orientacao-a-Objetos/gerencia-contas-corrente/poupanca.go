package main

import (
	"fmt"
	"learning-go/Go: Orientacao a Objetos/gerencia-contas-corrente/contas"
)

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDaPati := contas.ContaCorrente{}
	contaDoDenis.Depositar(5000)
	contaDoDenis.Sacar(500)

	fmt.Println(contaDoDenis)
	fmt.Println(contaDoDenis.ObterSaldo())

	fmt.Println(contaDaPati)

}

package main

import (
	"fmt"
	"learning-go/Go: Orientacao a Objetos/gerencia contas corrente/contas"
)

func main() {
	//clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor"}
	//contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100} //como o saque é um atributo privado, um erro será retornado
	//fmt.Println(contaDoBruno)

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())
}

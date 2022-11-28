package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "João"
	nome2 := "Ingrid" // Outra forma de declarar variavel e por padrao infere o tipo automaticamente
	var versao float32 = 1.1
	var idade int = 19
	var idade2 int //No GO se eu não atribuir um valor no tipo int a var recebe o valo 0 por padrão
	fmt.Println("Olá", nome2)
	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Olá, sr.", nome, "sua idade é", idade2)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
}

// Se eu declarar uma variavel e não utilizar, obrigatoriamente preciso retirar
//Não existe flout em go, exist flout32 e flout64

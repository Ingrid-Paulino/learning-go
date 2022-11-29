// Passando um valor ou uma cópia
package main

import (
	"fmt"
)

/*
Métodos são definidos de maneira parecida com funções, mas de uma maneira diferente. Existe um (p *Pessoa) que se refere a um ponteiro para a instância criada da estrutura, conforme o exemplo abaixo:
*/

type Pessoa struct {
	nome, sobrenome string
}

func (p *Pessoa) ExibirNomeCompleto() string {
	nomeCompleto := p.nome + " " + p.sobrenome
	return nomeCompleto
}

func main() {
	p1 := Pessoa{"Guilherme", "Lima"}
	fmt.Println(p1.ExibirNomeCompleto())
}

// temos a saída esperada:
// Guilherme Lima
// Nesse caso, passamos para o método o valor encontrado neste ponteiro através do (p *Pessoa).

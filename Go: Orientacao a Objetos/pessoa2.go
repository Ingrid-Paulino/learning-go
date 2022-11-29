// Passando uma cópia
package main

import (
	"fmt"
)

/*
Também é possível passar um valor removendo a assinatura do ponteiro (p *Pessoa) para (p Pessoa).

Nesse caso, uma cópia do valor de Pessoa é passada para a função, sem alterar o valor do ponteiro. Portanto, precisamos ficar atentos, já que qualquer alteração que você faça em p se passar por valor não será refletida na fonte p.
*/

type Pessoa2 struct {
	nome, sobrenome string
}

func (p Pessoa2) ExibirNomeCompleto2() string {
	p.sobrenome = "Silva"
	nomeCompleto := p.nome + " " + p.sobrenome
	return nomeCompleto
}

func main() {
	p1 := Pessoa2{"Guilherme", "Lima"}

	fmt.Println(p1.ExibirNomeCompleto2())
	fmt.Println(p1.nome, p1.sobrenome)
}

//Nossa saída será:
//Guilherme Silva
//Guilherme Lima

/*
Observe que alteramos o sobrenome de p no método ExibirNomeCompleto, mas não foi alterado o valor armazenado no ponteiro. Sendo assim, quando não precisamos alterar o conteúdo de um ponteiro, podemos passar apenas uma cópia.
*/

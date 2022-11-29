package contas

import "learning-go/Go: Orientacao a Objetos/gerencia contas corrente/clientes"

type ContaCorrente struct {
	//Para que os atributos fique visiveis para outros arquivos, é necessario declaralos com letra maiscula
	Titular       clientes.Titular //composição, no GO não existe herança //Composição de tipos é suportada e encorajada em Go.
	NumeroAgencia int
	NumeroConta   int
	saldo         float64 //esse atributo esta privado,
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

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

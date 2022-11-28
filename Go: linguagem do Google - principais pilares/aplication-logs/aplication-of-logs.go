/*
Essa aplicaçao monitora sites para ver se os mesmos estão online ou não
*/
package main

import (
	"bufio" //possibilita ler linha a linha do arquivo
	"fmt"
	"io"
	"io/ioutil"
	"net/http" //pacote que faz requisição web
	"os"       //conversa com o sistema operacional
	"strconv"  //converte varios tipos para string
	"strings"
	"time"
)

const monitoramentos = 3
const deley = 5

func main() {
	exibeIntroducao()
	//leSitesDoArquivo()
	//registraLog("dfdfdgfdg", false)
	for {
		exibeMenu()
		//var comando int
		//comando = leComando()
		//ou
		comando := leComando()

		switch comando { //No GO não tem break
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) //retorna status 0, o codigo 0 informo para o sistema operacional que ele saiu com sucesso
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) // -1 indico que teve algum problema para o sistema operacional na hora de sair o programa
		}
	}
}

func exibeIntroducao() {
	nome := "Ingrid"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int { //int é o tipo do retorno da func
	var comandoLido int
	// Capitura o que o usuario escreve
	fmt.Scanf("%d", &comandoLido)
	//ou
	//fmt.Scan(&comando) //nao preciso passar o modificador
	// %d sig que espero um inteiro
	//& sig que oq o usuario digitar eu quero salvar na variavel comando
	fmt.Println("O endereço da minha variavel comando é", &comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//sites := []string{"https://random-status-code.herokuapp.com/",
	//	"https://www.alura.com.br", "https://www.caelum.com.br"}

	sites := leSitesDoArquivo()

	//Primeira forma
	//site := "https://random-status-code.herokuapp.com/" //esse site retorna sempre um status code diferente
	//site := "https://www.alura.com.br"
	//resp, err := http.Get(site) //acessa o site //get retorna uma resposta=resp ou um erro=err
	//resp, _ := http.Get(site)
	//fmt.Println(resp)

	//Segunda forma
	for i := 0; i < monitoramentos; i++ { // testo 5 vezes
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(deley * time.Second) // Para cada vez que o primeiro for rodar, 5 segundos será esperado
		fmt.Println("")
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	basePath := "aplication-logs/sites.txt"
	//Forma 1 de ler arquivo
	//arquivo, err := os.Open(basePath) //abre um arquivo, e retorna de forma pura oq tem la dentro sem nenhuma conversão
	//Forma 2 de ler arquivo
	//arquivo, err := ioutil.ReadFile(basePath) // abre e retorna o arquivo em um array de bites, porem é mais facil de converter esse arquivo para algo que possamos entender

	arquivo, err := os.Open(basePath)

	//nil == null
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//Forma 2
	//fmt.Println(string(arquivo)) //converte a função ReadFile para algo entendivel // mostra o conteudo do arquivo de uma vez

	//Forma 3 de ler arquivo
	leitor := bufio.NewReader(arquivo) //possibilita ler linha a linha
	//leitor tem algumas funçoes que ajuda a passer pelo arquivo sites.txt
	//linha, err := leitor.ReadString('\n') //Ele vai ler ate a quebra de linha
	//if err != nil {
	//	fmt.Println("Ocorreu um erro:", err)
	//}
	//fmt.Println(linha) //retorna só a primeira linha

	//para o bufio ler todas as linhas
	for {
		linha, err := leitor.ReadString('\n') //Ele vai ler ate a quebra de linha
		linha = strings.TrimSpace(linha)      // tira os espaços entre cada linha retornada
		sites = append(sites, linha)
		//fmt.Println(linha)
		if err == io.EOF { //esse erro e lançado quando chega ao final do arquivo, pois nao tem mais linha para ler
			break // se chegr no final do arquivo, sai do loop for
		}

	}

	//fmt.Println(sites)
	arquivo.Close() //boa pratica //antes de eu encerrar, eu preciso fecha arquivo do sistema operacional para nao dar problema se outro programa quiser ler o mesmo arquivo
	return sites
}

func testaSite(site string) {
	//resp, _ := http.Get(site)
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

// restante do código omitido

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("aplication-logs/log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //OpenFile: se o arquivo nao existir ele cria // os.O_APPEND: para nao sobescrever o conteudo que ja existe

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + site + " - online: " + strconv.FormatBool(status) + "\n") //escreve no arquivo log.txt // strconv.FormatBool(status): converte booleano para string

	arquivo.Close()
	//fmt.Println(arquivo)
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("aplication-logs/log.txt") //Le e retorna tudo que tem no arquivo

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	// ioutil.ReadFile ja fecha o arquivo por isso nao é necessario fechar
	fmt.Println(string(arquivo))
}

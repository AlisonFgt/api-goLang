package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const _monitoramentos = 3
const _delay = 3

func main() {

	exibeIntroducao()

	for {
		// for sem parametro no golang é o while true

		nome, idade := devolveNomeEIdade()
		fmt.Println(nome, "e tenho ", idade, "anos")

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimiLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Alison"
	versao := 1.2
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O valor da variável comando é:", comandoLido)

	return comandoLido
}

func devolveNomeEIdade() (string, int) {
	nome := "Alison"
	idade := 31
	return nome, idade
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	//var sites [2]string array com posições fixas
	//sites[0] = "http://alisonalves.com"
	//sites[1] = "https://random-status-code.herokuapp.com"

	//sites := []string{"http://alisonalves.com", "https://random-status-code.herokuapp.com"} // slices com posições dinamicas conforme o uso

	sites := lerSitesDoArquivo()

	for i := 0; i < _monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":", site)
			testaSite(site)
		}
		time.Sleep(_delay * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func exibeNomes() {
	nomes := []string{"Alison", "Carlos", "Caio"}
	fmt.Println(nomes)
}

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha) // limpa string quebra linha ou espaço
		sites = append(sites, linha)     // atribui a um slices "array" o valor de um site

		if err == io.EOF {
			break // ao chegar ao final do arquivo sai do for == while
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	//https://golang.org/src/time/format.go pegar padrões de formatação de horas
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimiLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}

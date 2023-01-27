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

const const_MONITORAMENTO = 3
const const_DELAY = 5

func main() {

	ExibeIntroducao()

	for {

		ExibeMenu()

		switch leComando() {

		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do Programa ... ")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")

		}

	}

}

func ExibeIntroducao() {
	nome, _ := os.Hostname()
	versao := 1.2

	fmt.Println("Olá Usuário, ", nome)
	fmt.Println("versao do programa ", versao)
}

func leComando() int {

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando executado é: ", comando)

	return comando
}

func ExibeMenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando ...")

	sites := leArquivoTexto()

	for i := 0; i < const_MONITORAMENTO; i++ {

		for _, site := range sites {
			testaSites(site)
		}

		time.Sleep(const_DELAY * time.Second)
	}

}

func testaSites(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Aconteceu um erro :", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O ", site, "está online")
		registraLog(site, true)
	} else {
		fmt.Println("Ops o site: ", site, " está com algum problema:  Status Code: ", resp.StatusCode)
		registraLog(site, false)

	}

}

func leArquivoTexto() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Deu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Aconteceu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online " + strconv.FormatBool(status) + "\n")
	arquivo.Close()

}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}

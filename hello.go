package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	numerOfMonitoring = 5
	delay             = 5
	errorMessage      = "Ocorreu um erro:"
)

func main() {
	header()
	for {
		menu()
		command := readCommand()

		switch command {
		case 1:
			startsMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
			printLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não conhecido!")
			os.Exit(-1)
		}
	}
}

func header() {
	name := "Barriquero"
	version := 1.1
	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
	fmt.Println()
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
	fmt.Println()
}

func readCommand() int {
	var command int

	_, err := fmt.Scan(&command)
	if err != nil {
		log.Panic("Erro.", err)
	}

	fmt.Println("O comando escolhido foi", command)

	return command
}

func startsMonitoring() {
	fmt.Println("Monitorando...")

	sites := readFileWithSites()

	for i := 0; i < numerOfMonitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site:", i, ", Site:", site)
			siteTest(site)
		}

		fmt.Println()

		time.Sleep(delay * time.Second)
	}

	fmt.Println()
}

func siteTest(site string) {
	resp, err := http.Get(site)
	if err != nil {
		log.Panic(err)
	}

	statusCode := resp.StatusCode

	if statusCode >= 200 && statusCode <= 208 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		logRegister(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", statusCode)
		logRegister(site, false)
	}
}

func readFileWithSites() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(errorMessage, err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println(errorMessage, err)
	}

	return sites
}

func logRegister(site string, status bool) {
	file, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(errorMessage, err)
	}

	_, err = file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - site: " + site + " - online: " + strconv.FormatBool(status) + "\n")
	if err != nil {
		fmt.Println(errorMessage, err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println(errorMessage, err)
	}
}

func printLogs() {
	file, err := os.ReadFile("log.log")
	if err != nil {
		fmt.Println(errorMessage, err)
	}

	fmt.Println(string(file))
}

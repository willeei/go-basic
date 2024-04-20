package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	for {
		header()
		menu()
		command := readCommand()

		switch command {
		case 1:
			startsMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
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
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
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
	site := "https://httpstat.us/Random/200,201,208,400,404,500"
	resp, err := http.Get(site)
	if err != nil {
		log.Panic(err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 208 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoringTimes = 5
const delay = 5

func main() {

	showIntro()

	for {

		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Guilherme"
	version := 1.1
	fmt.Println("Hello, sr.", name)
	fmt.Println("This program is in version", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func startMonitoring() {
	fmt.Println("\nMonitoring...")

	//sites := []string{"https://www.alura.com.br", "https://httpstat.us/500", "https://www.caelum.com.br"}
	sites := readSitesFromFile()

	for i := 0; i < monitoringTimes; i++ {
		for _, site := range sites {
			testSite(site)
		}
		fmt.Println()
		time.Sleep(delay * time.Second)
	}

}

func testSite(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro ao fazer um get para o site", site, ":", err)
		createLog(site, false)
		return
	}
	if response.StatusCode == 200 {
		createLog(site, true)
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		createLog(site, false)
		fmt.Println("Site", site, "recebeu status", response.StatusCode, "e está com problemas!")
	}

}

func readSitesFromFile() []string {
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu o seguinte erro:", err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)
	sites := []string{}

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	file.Close()
	return sites

}

func createLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	currentTime := time.Now().Format("02/01/2006 15:04:05")
	file.WriteString(currentTime + "  " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	file, err := os.Open("log.txt")

	if err != nil {
		fmt.Println("Ocorreu o seguinte erro:", err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)
	fmt.Println()

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

}

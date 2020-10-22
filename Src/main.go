package main

import (
	"fmt"
	"time"

	"../Packages/extract"
)

func main() {

	now := time.Now()
	// Espera todos processos finalizarem
	defer func() {
		fmt.Println("\nTempo de execução:", time.Since(now))
	}()

	reader := extract.CsvReader()

	finished := make(chan bool)

	UFs := []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG", "PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}
	fmt.Println("Escolha de qual UF deseja extrair dados: ")

	for i := range UFs {
		fmt.Printf("%s ", UFs[i])
	}

	fmt.Print("\n->")

	var UF string
	fmt.Scan(&UF)

	if extract.Contains(UFs, UF) == true {
		go extract.MeanScoresUF(reader, UF, finished)

		// read channel
		<-finished
	}

	return
}

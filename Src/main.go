package main

import (
	"fmt"
	"time"

	"../Packages/extract"
)

func main() {

	// cria um canal para goroutine
	finished := make(chan bool)	
	
	reader := extract.CsvReader()

	siglas := []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG", 
		"PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}

	// gera array de estruturas de Estado (UFs)
	states := []extract.State{}
	for i := range siglas {
		states = append(states, extract.NewState(siglas[i]))
	}

	opcao := 0
	fmt.Println("Escolha uma opção:")
	fmt.Printf("\t1) Sem go routines\n")
	fmt.Printf("\t2) Com go routines\n")
	fmt.Print("\n-> ")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		now := time.Now()
		// defer - Espera todos processos finalizarem
		defer func() {
			fmt.Println("\n\nTempo de execução:", time.Since(now))
		}()
		
		extract.UFData(reader, states)
	
		break
	case 2:
		now := time.Now()
		defer func() {
			fmt.Println("\n\nTempo de execução:", time.Since(now))
		}()

		extract.UFDataPallel(reader, states)

		// recebe 
		<-finished
		
		break
	default:
		fmt.Println("Opção Inválida!")
		break;
	}
		
	
	return
}

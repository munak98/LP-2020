package main

import (
	"fmt"
	"time"

	"../Packages/extract"
)

func main() {

	finished := make(chan bool)	// cria um canal para mensurar tempos
	
	reader := extract.CsvReader()

	UFs := []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG", 
		"PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}
	fmt.Println("Escolha de qual UF deseja extrair dados: ")
	for i := range UFs {
		fmt.Printf("%s ", UFs[i])
	}
	fmt.Print("\n-> ")

	var UF string
	fmt.Scan(&UF)

	if extract.Contains(UFs, UF) == true {
		
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
			
			extract.MeanScoresUF2(reader, UF)
		
			break
		case 2:
			now := time.Now()
			defer func() {
				fmt.Println("\n\nTempo de execução:", time.Since(now))
			}()

			go extract.MeanScoresUF(reader, UF, finished)

			//leitura do canal
			<-finished
			break
		default:
			fmt.Println("Opção Inválida!")
			break;
		}
	} else {
		fmt.Print("Escolha inválida!")
	}
	
	return
}

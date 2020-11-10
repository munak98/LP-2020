package main

import (
	"fmt"

	"../Packages/extract"
)

func main() {

	// cria um canal para goroutine
	finished := make(chan bool)	
	
	reader := extract.CsvReader()

	states := extract.NewStates();

	opcao := 0
	fmt.Println("Escolha uma opção:")
	fmt.Printf("\t1) Sem go routines\n")
	fmt.Printf("\t2) Com go routines\n")
	fmt.Print("\n-> ")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		states = extract.Data(reader, states)

		break
	case 2:

		// cria um canal para goroutine
		//statesChannel := make(chan []extract.State)

		go extract.DataPallel(reader, &states, finished)

		// recebe 
		<-finished
		
		break
	default:
		fmt.Println("Opção Inválida!")
		break;
	}

	extract.Menu(states)

	return
}

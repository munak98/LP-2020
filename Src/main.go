package main

import (
	"fmt"

	"../Packages/extract"
)

func main() {

	states := extract.NewStates()

	opcao := 0
	fmt.Println("Escolha uma opção:")
	fmt.Printf("\t1) Sem go routines\n")
	fmt.Printf("\t2) Com go routines\n")
	fmt.Print("\n-> ")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:

		states = extract.Data(states)

		break
	case 2:

		extract.DataPallel(&states)
		
		break
	default:
		fmt.Println("Opção Inválida!")
		break;
	}

	extract.Menu(states)

	return
}

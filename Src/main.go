package main

import (
	"fmt"
	"os"
	"runtime"

	"../Packages/extract"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	states := extract.NewStates()

	opcao := 0
	fmt.Println("Escolha uma opção para extrair dados:")
	fmt.Printf("Digite 0 para sair\n\n")
	fmt.Printf("\n\t1) Sem go routines\n")
	fmt.Printf("\t2) Com go routines\n")
	fmt.Print("\n-> ")
	fmt.Scan(&opcao)

	switch opcao {
	case 0:

		os.Exit(3)
		break
	case 1:

		states = extract.Data(states)
		break
	case 2:

		extract.DataParallel(&states)
		break
	default:
		fmt.Println("Opção Inválida!")
		break
	}

	extract.MostParticipantsUF(states)

	extract.Menu(states)

	return
}

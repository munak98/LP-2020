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
//em go as variaveis sempre são iniciadas em zero, então neste caso podemos colocar "var opcao int" 
	opcao := 0
	fmt.Println("Escolha uma opção para extrair dados:")
	fmt.Printf("Digite 0 para sair\n\n")
	fmt.Printf("\n\t1) Sem Paralelismo\n")
	fmt.Printf("\t2) Com Paralelismo\n")
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
		os.Exit(3)
		break
	}

	extract.MostParticipantsUF(states)

	extract.Menu(states)

	return
}

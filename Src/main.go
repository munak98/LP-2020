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

	years := extract.NewYears()

	// extract.FileInfo(years[0].CsvFilePath)

	//em go as variaveis sempre são iniciadas em zero, então neste caso podemos colocar "var opcao int" 
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

		years = extract.Data(years)
		break
	case 2:

		extract.DataParallel(&years)
		break

	case 3:

		contents := extract.GetFilesContents(years[0].CsvFilePath, years[1].CsvFilePath)
		fmt.Println(contents)

		break
	default:
		fmt.Println("Opção Inválida!")
		break
	}

	for i := range years {
		extract.MostParticipantsUF(years[i].States)
	}

	extract.Menu(years[2].States)

	return
}

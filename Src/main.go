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

	var opcao int
	fmt.Println("Escolha uma opção para extrair dados:")
	fmt.Println("Digite -1 para sair")
	fmt.Println("\n1) Sem Paralelismo")
	fmt.Println("2) Com Paralelismo")
	fmt.Print("\n-> ")
	fmt.Scan(&opcao)

	switch opcao {
	case -1:
		os.Exit(3)
	case 1:
		years = extract.Data(years)
	case 2:
		extract.DataParallel(&years)
	case 3:
		contents := extract.GetFilesContents(years[0].CsvFilePath, years[1].CsvFilePath)
		fmt.Println(contents)
	default:
		fmt.Println("Opção Inválida!")
		os.Exit(3)
	}

	extract.YearsInfo(years)
	extract.YearsMenu(years)

	return
}

package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
	"../Packages/extract"
)

type SchoolScores = extract.SchoolScores
type RaceScores = extract.RaceScores

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	csvFilePath19 := "../microdados_enem_2019/DADOS/MICRODADOS_ENEM_2019.csv"
	csvFilePath18 := "../microdados_enem_2018/DADOS/MICRODADOS_ENEM_2018.csv"
	csvFilePath17 := "../microdados_enem_2017/DADOS/MICRODADOS_ENEM_2017.csv"
	states19 := extract.NewStates()
	states18 := extract.NewStates()
	states17 := extract.NewStates()

	schoolScores := [3]SchoolScores{}
	raceScores := [3]RaceScores{}
	count := 0

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
		fmt.Println("Extraindo dados..")
		now := time.Now()
		states19 = extract.Data(states19, &schoolScores, &raceScores, csvFilePath19, &count)
		states18 = extract.Data(states18, &schoolScores, &raceScores, csvFilePath18, &count)
		states17 = extract.Data(states17, &schoolScores, &raceScores, csvFilePath17, &count)
		fmt.Println("\n\nTempo de execução:", time.Since(now))
		break
	case 2:
		fmt.Println("Extraindo dados..")
		now := time.Now()
		extract.DataParallel(&states19, &schoolScores, &raceScores, csvFilePath19, &count)
		extract.DataParallel(&states18, &schoolScores, &raceScores, csvFilePath18, &count)
		extract.DataParallel(&states17, &schoolScores, &raceScores, csvFilePath17, &count)
		fmt.Println("\n\nTempo de execução:", time.Since(now))
		break
	default:
		fmt.Println("Opção Inválida!")
		break
	}

	// extract.MostParticipantsUF(states19)
	fmt.Println("Número de registros analisados:\n\n", count)
	extract.YearsMenu(states19, states18, states17, &schoolScores, &raceScores)

	return
}

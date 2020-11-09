package extract

import (
	"encoding/csv"
	"fmt"
	"io"
)

/* SEM PARALELISMO */

//UFData pega os dados de Medias das notas de uma UF do arquivo CSV - Sem Goroutine!
func UFData(reader *csv.Reader, states []State) {

	// notas de areas de conhecimento da UF
	scoresUF := [4][]float64{}

	// notas de areas de conhecimento de cada raça
	scoresPerRace := [6][4][]float64{}

	count := 0

	// leitura de linha a linha do registro
	for /* i := 0; i < 1000000; i++ */ {
		recordLine, err := reader.Read()

		if err == io.EOF {
			break // chegou ao final do registro
		} else if err != nil { //checa por outros erros
			fmt.Println("An error encountered ::", err)
		}

		count++

		for i := range states {
			if states[i].Sigla == recordLine[5] {

				states[i].Total++

				// coleta as notas de cada disciplina de toda UF
				getScores(recordLine, &scoresUF)

				// coleta dados por raça da UF
				getRacesData(recordLine, &states[i], &scoresPerRace)
			}
		}

	}

	fmt.Println("Numero de registros analisados:", count)

	for i := range states {

		states[i].Medias = getMeanScores(scoresUF)

		for j := range states[i].Races {
			states[i].Races[j].Medias = getMeanScores(scoresPerRace[j])
		}

		// printUFMeanScores(states[i])
		// printRacesMeanScores(states[i])
	}

	printUFMeanScores(states[6])
	printRacesMeanScores(states[6])

	return
}

/* COM PARALELISMO */

//UFData pega os dados de Medias das notas de uma UF do arquivo CSV
func UFDataPallel(reader *csv.Reader, states []State) {

	// notas de areas de conhecimento da UF
	scoresUF := [4][]float64{}

	// notas de areas de conhecimento de cada raça
	scoresPerRace := [6][4][]float64{}

	count := 0

	// ch := make(chan bool)

	// var wg sync.WaitGroup

	// leitura de linha a linha do registro
	for /* i := 0; i < 1000000; i++ */ {
		recordLine, err := reader.Read()
		count++

		if err == io.EOF {
			break // chegou ao final do registro
		} else if err != nil { //checa por outros erros
			fmt.Println("An error encountered ::", err)
		}

		for i := range states {
			if states[i].Sigla == recordLine[5] {

        // wg.Add(1)

        states[i].Total++

        // coleta as notas de cada disciplina de toda UF
        go getScores(recordLine, &scoresUF)

        // coleta dados por raça da UF
        go getRacesData(recordLine, &states[i], &scoresPerRace)

			}
		}
  }
  
  // go func() {
  //   wg.Wait()
  //   close(ch)
  // }()

	fmt.Println("Numero de registros analisados:", count)

	for i := range states {

		states[i].Medias = getMeanScores(scoresUF)

		for j := range states[i].Races {
			states[i].Races[j].Medias = getMeanScores(scoresPerRace[j])
		}

		// printUFMeanScores(states[i])
		// printRacesMeanScores(states[i])
	}

	go printUFMeanScores(states[6])
	go printRacesMeanScores(states[6])

	return
}

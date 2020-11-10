package extract

import (
	"encoding/csv"
	"fmt"
	"io"
	"sync"
	"time"
)

/* SEM PARALELISMO */

//Data pega os dados de Medias das notas de todos Estados do arquivo CSV
func Data(reader *csv.Reader, states []State) []State {

	now := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(now))
	}()

	count := 0

	// leitura de linha a linha do registro
	for /* i := 0; i < 1000; i++ */ {
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

				// coleta as notas de cada area da UF
				getUFScores(recordLine, &states[i])

				// coleta dados de cada area por raça da UF
				getRacesData(recordLine, &states[i])
			}
		}
		// fmt.Println("Processando linha:", count)

	}
	fmt.Println("Numero de registros analisados:", count)

	for i := range states {

		states[i].Medias = getMeanScores(states[i].Scores)

		for j := range states[i].Races {
			states[i].Races[j].Medias = getMeanScores(states[i].Races[j].Scores)
		}

		// Printa todas UFs
		// PrintUFMeanScores(states[i])
		// PrintRacesMeanScores(states[i])
	}

	// printa apenas do DF
	// PrintUFMeanScores(states[6])
	// PrintRacesMeanScores(states[6])

	return states
}

/* COM PARALELISMO */

//DataPallel pega os dados de Medias das notas de uma UF do arquivo CSV
func DataPallel(reader *csv.Reader, states []State) {

	now := time.Now()
	// defer - Espera todos processos finalizarem
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(now))
	}()

	count := 0

	ch := make(chan bool)

	var wg sync.WaitGroup

	go func() {
		// leitura de linha a linha do registro
		for i := 0; i < 50; i++ {
			recordLine, err := reader.Read()

			if err == io.EOF {
				break // chegou ao final do registro
			} else if err != nil { //checa por outros erros
				fmt.Println("An error encountered ::", err)
			}

			count++
			defer wg.Done()
			wg.Add(1)

			for i := range states {
				if states[i].Sigla == recordLine[5] {

					states[i].Total++

					// coleta as notas de cada area da UF
					getUFScores(recordLine, &states[i])

					// coleta dados de cada area por raça da UF
					getRacesData(recordLine, &states[i])

					fmt.Println("Processando linha:", count)
				}
			}
			ch <- true
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("Numero de registros analisados:", count)

	for i := range states {

		states[i].Medias = getMeanScores(states[i].Scores)

		for j := range states[i].Races {
			states[i].Races[j].Medias = getMeanScores(states[i].Races[j].Scores)
		}

		// Printa todas UFs
		// PrintUFMeanScores(states[i])
		// PrintRacesMeanScores(states[i])
	}

	// printa apenas do DF
	PrintUFMeanScores(states[6])
	PrintRacesMeanScores(states[6])

	return
}

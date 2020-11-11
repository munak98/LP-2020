package extract

import (
	"fmt"
	"io"
	"sync"
	"time"
)

/* SEM PARALELISMO */

//Data pega os dados de todos Estados do arquivo CSV
func Data(states []State) []State {

	now := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(now))
	}()

	reader := CsvReader()
	count := 0

	// leitura de linha a linha do registro
	for /* i := 0; i < 500000; i++ */ {
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
	fmt.Println("\nNumero de registros analisados:", count)

	getStatesMeanScores(&states)

	return states
}

/* COM PARALELISMO */

//DataParallel pega os dados de forma paralela de todos Estados do arquivo CSV
func DataParallel(states *[]State) *[]State {

	start := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(start))
	}()

	reader := CsvReader()
	var wg sync.WaitGroup
	count := 0

	// leitura de linha a linha do registro
	for /* i := 0; i < 500000; i++ */ {
		recordLine, err := reader.Read()

		if err == io.EOF {
			break // chegou ao final do registro
		} else if err != nil { //checa por outros erros
			fmt.Println("An error encountered ::", err)
		}
		count++

		go func() {
			wg.Add(1)
			defer wg.Done()

			for i := range *states {
				if (*states)[i].Sigla == recordLine[5] {

					(*states)[i].Total++
	
					// coleta as notas de cada area da UF
					getUFScores(recordLine, &(*states)[i])
	
					// coleta dados de cada area por raça da UF
					getRacesData(recordLine, &(*states)[i])
				}
			}
			//fmt.Println("Processando linha:", count)
		}()
	}

	wg.Wait()
	
	fmt.Println("Numero de registros analisados:", count)

	getStatesMeanScores(states)

	return states
}


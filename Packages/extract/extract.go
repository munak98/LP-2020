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
	for i := 0; i < 500000; i++ {
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

//DataPallel pega os dados de todos Estados do arquivo CSV
func DataPallel(states *[]State) {

	start := time.Now()
	// defer - Espera todos processos finalizarem
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(start))
	}()

	reader := CsvReader()
	ch := make(chan []string)
	var wg sync.WaitGroup
	count := 0

	// leitura de linha a linha do registro
	for  i := 0; i < 500000; i++  {
		recordLine, err := reader.Read()

		if err == io.EOF {
			break // chegou ao final do registro
		} else if err != nil { //checa por outros erros
			fmt.Println("An error encountered ::", err)
		}
		count++
		wg.Add(1)

		go func(record []string, states *[]State, count int) {
			defer wg.Done()
			for i := range *states {

				if (*states)[i].Sigla == recordLine[5] {

					(*states)[i].Total++

					// coleta as notas de cada area da UF
					getUFScores(recordLine, &(*states)[i])

					// coleta dados de cada area por raça da UF
					getRacesData(recordLine, &(*states)[i])

					//fmt.Println("Processando linha:", count)
				}
			}
			ch <- record
		}(recordLine, states, count)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// print channel results (necessary to prevent exit programm before)
	j := 0
	for range ch {
		j++
		//fmt.Printf("\r\t\t\t\t | done %d\n", j)
	}

	fmt.Println("Numero de registros analisados:", count)

	getStatesMeanScores(states)

	//finished <- true

	return
}

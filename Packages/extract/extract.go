package extract

import (
	"fmt"
	"io"
	"sync"
	"time"
)

/* SEM PARALELISMO */

//Data pega os dados de todos Estados do arquivo CSV
func Data(years []Year) []Year {

	now := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(now))
	}()

	fmt.Println("Extraindo dados..")

	// loop pelos anos
	for i := range years {

		reader, _ := CsvReader(years[i].CsvFilePath)
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
			
			// loop pelos estados
			for j := range years[i].States {
				if years[i].States[j].Sigla == recordLine[5] {

					years[i].States[j].Total++

					getYearSchoolScores(recordLine, &years[i])

					// Pega as notas estaduais de areas de conhecimento
					getScores(recordLine, &years[i].States[j].Scores, years[i].Year)

					// coleta dados de cada area por raça da UF
					getRacesData(recordLine, &years[i].States[j], years[i].Year)
				}
			}

			//fmt.Println("Processando linha:", count)
		}

		fmt.Printf("Numero de registros analisados de %d: %d\n", years[i].Year, count)
		count = 0 // Reseta contagem a cada ano

		getStatesMeanScores(&years[i].States)
	}

	getYearsMeanScores(&years)

	return years
}

/* COM PARALELISMO */

//DataParallel pega os dados de todos Estados do arquivo CSV
func DataParallel(years *[]Year) {
	start := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(start))
	}()
	fmt.Println("Extraindo dados..")
	var wg sync.WaitGroup

	for i := range *years {
		reader, _ := CsvReader((*years)[i].CsvFilePath)
		wg.Add(1)

		totalRecords := (*years)[i].Total
		workers := (*years)[i].Workers

		// Execução paralela dos processos
		go func(i int) {
			defer wg.Done()
			// Loop no numero de processos necessários para cada ano
			for j := 0; j < (*years)[i].Workers; j++ {
				getData(
					reader,
					&(*years)[i],
					(totalRecords/workers)*j,     // inicio do pedaço
					(totalRecords/workers)*(j+1), // final do pedaço
				)
			}
		}(i)
	}
	wg.Wait()

	for i := range *years {
		getStatesMeanScores(&(*years)[i].States)
		fmt.Printf("Numero de registros analisados de %d: %d\n", (*years)[i].Year, (*years)[i].Total)
	}

	getYearsMeanScores(&(*years))
}

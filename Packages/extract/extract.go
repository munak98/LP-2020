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

					// coleta as notas de cada area da UF
					getUFScores(recordLine, &years[i].States[j])

					// coleta dados de cada area por raça da UF
					getRacesData(recordLine, &years[i].States[j])
				}
			}

			//fmt.Println("Processando linha:", count)
		}

		fmt.Printf("Numero de registros analisados de %d: %d", years[i].Year, count)
		count = 0 // Reseta contagem a cada ano
	}

	for i := range years {
		getStatesMeanScores(&years[i].States)
	}

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
	count := 0

	for i := range *years {
		reader, _ := CsvReader((*years)[i].CsvFilePath)
		wg.Add(1)

		// Execução paralela dos processos
		go func(i int) {	
			defer wg.Done()

			// Loop no numero de processos necessários para cada ano
			for j := 0; j < (*years)[i].Workers; j++ {
				getData(
					reader,
					&(*years)[i].States,
					&count,
					((*years)[i].TotalRecords/(*years)[i].Workers)*j,			// inicio do pedaço
					((*years)[i].TotalRecords/(*years)[i].Workers)*(j+1),	// final do pedaço
				)
			}
		}(i)
		count = 0
	}

	wg.Wait()

	for i := range *years {
		getStatesMeanScores(&(*years)[i].States)
	}

	return
}

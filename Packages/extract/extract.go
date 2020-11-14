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
		fmt.Println("\nNumero de registros analisados:", count)
		count = 0	// Reseta contagem a cada ano
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

	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := range *years {
			
			reader, _ := CsvReader((*years)[i].CsvFilePath)
			
			for j := 0; i < (*years)[i].Workers; j++ {
				getData(
					reader, 
					&(*years)[i].States, 
					&count, 
					(*years)[i].TotalRecords / (*years)[i].Workers*j, 
					(*years)[i].TotalRecords / (*years)[i].Workers*(j+1),
				)
			}
			
			count = 0
		}
	}()
		
	wg.Wait()
	fmt.Println("Numero de registros analisados:", count)

	for i := range *years {
		getStatesMeanScores(&(*years)[i].States)
	}

	return
}

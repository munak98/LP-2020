package extract

import (
	"fmt"
	"io"
	"time"
)

/* SEM PARALELISMO */

//Data pega os dados de todos Estados do arquivo CSV
func Data(states []State) []State {

	now := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(now))
	}()

	fmt.Println("Extraindo dados..")

	reader, _ := CsvReader()
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
		//fmt.Println("Processando linha:", count)

	}
	fmt.Println("\nNumero de registros analisados:", count)

	getStatesMeanScores(&states)

	return states
}

/* COM PARALELISMO */

//DataParallel pega os dados de todos Estados do arquivo CSV
func DataParallel(states *[]State) {

	start := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(start))
	}()

	fmt.Println("Extraindo dados..")

	count := 0
	reader, _ := CsvReader()

	// fileSize := int(fileInfo.Size())
	// fmt.Println("Filinfo size: ", fileSize)

	//* total de registros
	const totalRecords = 5095271
	divisor := 29

	for i := 0; i < divisor; i++ {
		getData(reader, states, &count, totalRecords/divisor*i, totalRecords/divisor*(i+1))
	}
	
	fmt.Println("Numero de registros analisados:", count)

	getStatesMeanScores(states)

	return
}


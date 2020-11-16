package concurrency_extract

import (
	"fmt"
	"io"
	"sync"
)

/* SEM PARALELISMO */

//Data pega os dados de todos Estados do arquivo CSV
func Data(states []State, school *[3]SchoolScores, race *[3]RaceScores, csvFilePath string, count *int) []State {
	wg := new(sync.WaitGroup)

	reader, _ := CsvReader(csvFilePath)

	// leitura de linha a linha do registro
	for {
		recordLine, err := reader.Read()

		if err == io.EOF {
			break // chegou ao final do registro
		} else if err != nil { //checa por outros erros
			fmt.Println("An error encountered ::", err)
		}
		*count++
		var year int
		if csvFilePath == "../microdados_enem_2019/DADOS/MICRODADOS_ENEM_2019.csv"{
			year = 2
			wg.Add(1)
			go getGeneralRaceScores(recordLine, &(*race)[2], year, wg)
			wg.Add(1)
			go getSchoolScores(recordLine, &(*school)[2], year, wg)
		}
		if csvFilePath == "../microdados_enem_2018/DADOS/MICRODADOS_ENEM_2018.csv"{
			year = 1
			wg.Add(1)
			go getGeneralRaceScores(recordLine, &(*race)[1], year, wg)
			wg.Add(1)
			go getSchoolScores(recordLine, &(*school)[1], year, wg)
		}
		if csvFilePath == "../microdados_enem_2017/DADOS/MICRODADOS_ENEM_2017.csv"{
			year = 0
			wg.Add(1)
			go getGeneralRaceScores(recordLine, &(*race)[0], year, wg)
			wg.Add(1)
			go getSchoolScores(recordLine, &(*school)[0], year, wg)
		}



		for i := range states {
			if states[i].Sigla == recordLine[5] {

				states[i].Total++

				// coleta as notas de cada area da UF
				wg.Add(1)
				go getUFScores(recordLine, &states[i], year, wg)

				// coleta dados de cada area por raça da UF
				wg.Add(1)
				go getRacesData(recordLine, &states[i], year, wg)

			}
		}
		wg.Wait()

	}

	// fmt.Println("\nNúmero de registros analisados:", count)

	getStatesMeanScores(&states)

	return states
}

/* COM PARALELISMO */

//DataParallel pega os dados de todos Estados do arquivo CSV
func DataParallel(states *[]State, school *[3]SchoolScores, race *[3]RaceScores, csvFilePath string, count *int) {


	reader, _ := CsvReader(csvFilePath)

	// fileSize := int(fileInfo.Size())
	// fmt.Println("Filinfo size: ", fileSize)
	var year int

	//* total de registros
	var totalRecords int
	if csvFilePath == "../microdados_enem_2019/DADOS/MICRODADOS_ENEM_2019.csv"{
		totalRecords = 5095271
		year = 2
	}
	if csvFilePath == "../microdados_enem_2018/DADOS/MICRODADOS_ENEM_2018.csv"{
		totalRecords = 5513748
		year = 1
	}
	if csvFilePath == "../microdados_enem_2017/DADOS/MICRODADOS_ENEM_2017.csv"{
		totalRecords = 6731342
		year = 0
	}
	divisor := 29

	for i := 0; i < divisor; i++ {
		getData(reader, states, &(*school)[year], &(*race)[year], count, totalRecords/divisor*i, totalRecords/divisor*(i+1), year)
	}

	// fmt.Println("Número de registros analisados:", count)

	getStatesMeanScores(states)

	return
}

package extract

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/text/encoding/charmap"
)

// CsvReader faz a leitura do arquivo csv
func CsvReader() *csv.Reader {

	csvFilePath := "../microdados_enem_2019/DADOS/MICRODADOS_ENEM_2019.csv"
	fmt.Printf("FilePath: %s\n", csvFilePath)

	csvFile, err := os.Open(csvFilePath)

	if err != nil { // chaca se ocorre erros na abertura do csv
		fmt.Println("An error encountered ::", err)
	}

	reader := csv.NewReader(charmap.ISO8859_1.NewDecoder().Reader(csvFile))
	reader.Comma = ';'

	return reader
}

// Contains verifica se existe string em um array de strings
func Contains(list []string, str string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}

//N funfa..
func MeasureTime() {
	now := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(now))
	}()
}

// printa dados acerca da UF
func printUFMeanScores(state State) {
	fmt.Println("---------------------------------------")
	fmt.Printf("\nDados de %s", state.UF)
	fmt.Printf("\nTotal de participantes: %d\n", state.Total)
	fmt.Println("Médias:")

	fmt.Printf("\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n",
		state.Medias[0],
		state.Medias[1],
		state.Medias[2],
		state.Medias[3],
	)

	fmt.Printf("Numero de participantes de Escola Publica: %d\n", state.SchoolType[1])
	fmt.Printf("Numero de participantes de Escola Privada: %d\n", state.SchoolType[2])
}

// printa dados acerca de cada raça
func printRacesMeanScores(state State) {

	for i := range state.Races {
		fmt.Println("---------------------------------------")
		fmt.Printf("\nDados de %s", state.Races[i].Name)
		fmt.Printf("\nTotal de participantes: %d\n", state.Races[i].Total)
		fmt.Printf("Médias: \n\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n",
			state.Races[i].Medias[0],
			state.Races[i].Medias[1],
			state.Races[i].Medias[2],
			state.Races[i].Medias[3],
		)
		fmt.Printf("Numero de participantes de Escola Publica: %d\n", state.Races[i].SchoolType[1])
		fmt.Printf("Numero de participantes de Escola Privada: %d\n", state.Races[i].SchoolType[2])
	}

}

// pega o total de registros do arquivo CSV
func totalRecords(reader *csv.Reader) int {
	count := 0
	// leitura de linha a linha do registro
	for {
		_, err := reader.Read()

		if err == io.EOF {
			break // chegou ao final do registro
		} else if err != nil { //checa por outros erros
			fmt.Println("An error encountered ::", err)
		}
		count++
	}
	return count
}

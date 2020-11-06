package extract

import (
	"encoding/csv"
	"fmt"
	"os"
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

func printUFMeanScores(state State) {
	fmt.Println("---------------------------------------")
	fmt.Printf("\nDados de %s", state.uf)
	fmt.Printf("\nTotal de participantes: %d\n", state.totalParticipants)
	fmt.Println("Médias:")
	
	fmt.Printf("\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n",
		state.medias[0], 
		state.medias[1], 
		state.medias[2], 
		state.medias[3], 
	)
	
}
func printRacesMeanScores(state State) {

	for i := range state.races {
		fmt.Println("---------------------------------------")
		fmt.Printf("\nDados de %s", state.races[i].name)
		fmt.Printf("\nTotal de participantes: %d\n", state.races[i].total)
		fmt.Printf("Médias: \n\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n",
			state.races[i].medias[0], 
			state.races[i].medias[1], 
			state.races[i].medias[2], 
			state.races[i].medias[3], 
		)
	}

}


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

func printMeanScores(topico string, meanScores [4]float64) {
	fmt.Printf("Médias de %s: \n\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n", 
		topico, meanScores[0], meanScores[1], meanScores[2], meanScores[3])
}

// Contains verifica se existe string em um array de strings
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
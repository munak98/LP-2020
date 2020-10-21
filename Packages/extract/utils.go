package extract

import (
	"encoding/csv"
	"fmt"
	"os"
	"golang.org/x/text/encoding/charmap"
)

//CsvReader Reads the csv File
func CsvReader() *csv.Reader {

	csvFilePath := "../microdados_enem_2019/DADOS/MICRODADOS_ENEM_2019.csv"
	fmt.Printf("FilePath: %s\n", csvFilePath)

	csvFile, err := os.Open(csvFilePath)

	if err != nil { //check for error in opening
		fmt.Println("An error encountered ::", err)
	}

	reader := csv.NewReader(charmap.ISO8859_1.NewDecoder().Reader(csvFile))
	reader.Comma = ';'

	return reader
}

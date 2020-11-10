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

// Contains verifica se existe string na propiedade sigla de states
func Contains(states []State, str string) bool {

	for _, item := range states {
		if item.Sigla == str {
			return true
		}
	}

	return false
}

//MeasureTime N funfa..
func MeasureTime() {
	now := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(now))
	}()
}

//Menu para escolher quais dados mostrar
func Menu(states []State) {
	UFs := []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG", 
		"PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}
	fmt.Println("Escolha de qual UF deseja extrair dados: ")
	for i := range UFs {
		fmt.Printf("%s ", UFs[i])
	}
	fmt.Print("\n-> ")

	var UF string
	fmt.Scan(&UF)

	// verifica se existe UF no arrays de structs states
	if Contains(states, UF) == true {

		for i := range states {
			if UF == states[i].Sigla{
				PrintUFMeanScores(states[i])
				PrintRacesMeanScores(states[i])
			}
		}

	}else {
		fmt.Print("UF digitada inválida!")
	}
}

// printa dados acerca da UF
func PrintUFMeanScores(state State) {
	fmt.Println("---------------------------------------")
	fmt.Printf("\nDados de %s", state.Sigla)
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
func PrintRacesMeanScores(state State) {

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
// cuidado ao usar pois ao percorrer todo CSV, 
// tem que reler para percorrer de novo
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

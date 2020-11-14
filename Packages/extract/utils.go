package extract

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"golang.org/x/text/encoding/charmap"
)

// CsvReader faz a leitura do arquivo csv
func CsvReader(csvFilePath string) (*csv.Reader, os.FileInfo) {

	fmt.Printf("FilePath: %s\n", csvFilePath)

	csvFile, err := os.Open(csvFilePath)
	if err != nil { // checa se ocorre erros na abertura do csv
		fmt.Println("An error encountered ::", err)
	}

	fileInfo, err := csvFile.Stat()
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	reader := csv.NewReader(charmap.ISO8859_1.NewDecoder().Reader(csvFile))
	reader.Comma = ';'

	return reader, fileInfo
}

//GetFilesContents le varios arquivos de forma concorrente e retorna em um map
func GetFilesContents(files ...string) map[string][]byte {
	var wg sync.WaitGroup
	var m sync.Mutex

	filesLength := len(files)
	contents := make(map[string][]byte, filesLength)
	wg.Add(filesLength)

	for _, file := range files {
		go func(file string) {
			content, err := ioutil.ReadFile(file)

			if err != nil {
				log.Fatal(err)
			}

			m.Lock()
			contents[file] = content
			m.Unlock()
			wg.Done()
		}(file)
	}

	wg.Wait()

	return contents
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

	for {
		fmt.Println("\n\nEscolha de qual UF deseja visualizar dados: ")
		fmt.Printf("Digite 0 para sair\n\n")
		for i := range states {
			fmt.Printf("%s ", states[i].Sigla)
		}
		fmt.Print("\n-> ")

		var UF string
		fmt.Scan(&UF)

		// verifica se existe UF no arrays de structs states
		if Contains(states, UF) == true {

			for i := range states {
				if UF == states[i].Sigla {
					PrintUFMeanScores(states[i])
					PrintRacesMeanScores(states[i])
				}
			}

		} else if UF == "0" {
			break
		} else {
			fmt.Print("UF digitada inválida!")
		}

		UF = ""

		// tentativa de clear screen
		c := exec.Command("cls")
		c.Stdout = os.Stdout
		c.Run()
	}
}

//FileInfo - Pega o total de registros e tamanho do arquivo CSV
func FileInfo(csvFilePath string) int {

	reader, fileInfo := CsvReader(csvFilePath)
	total := 0

	fileSize := int(fileInfo.Size())
	fmt.Println("Filinfo size: ", fileSize)

	// leitura de linha a linha do registro
	for {
		_, err := reader.Read()

		if err == io.EOF {
			break // chegou ao final do registro
		} else if err != nil { //checa por outros erros
			fmt.Println("An error encountered ::", err)
		}
		total++
	}

	fmt.Println("Total de Registros: ", total)

	return total
}

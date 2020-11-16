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

//YearsMenu - para escolher quais dados mostrar sobre anos
func YearsMenu(years []Year) {

	for {
		fmt.Printf("\n*****************************************\n")
		fmt.Println("Escolha o ano.")
		fmt.Printf("Digite -1 para sair\n\n")
		fmt.Printf("0-Enem 2017\n1-Enem 2018\n2-Enem 2019\n")
		fmt.Print("\n-> ")

		var year int
		fmt.Scan(&year)

		fmt.Printf("\n*****************************************\n")
		// verifica se existe UF no arrays de structs states
		switch year {
		case -1:
			os.Exit(3)
		case 0:
			fmt.Printf("Dados nacionais Enem 2017\n")
			PrintYearMeanScores(years[year])
			PrintYearRacesMeanScores(years[year])
			PrintYearSchoolsMeanScores(years[year])
			MenuStates(years[year].States)
		case 1:
			fmt.Printf("Dados nacionais Enem 2018\n")
			PrintYearMeanScores(years[year])
			PrintYearRacesMeanScores(years[year])
			PrintYearSchoolsMeanScores(years[year])
			MenuStates(years[year].States)
		case 2:
			fmt.Printf("Dados nacionais Enem 2019\n")
			PrintYearMeanScores(years[year])
			PrintYearRacesMeanScores(years[year])
			PrintYearSchoolsMeanScores(years[year])
			MenuStates(years[year].States)
		default:
			fmt.Println("Opção inválida!")
			os.Exit(3)
		}
		fmt.Printf("\n*****************************************\n")
	}
}

//MenuStates para escolher quais dados mostrar sobre estados
func MenuStates(states []State) {

	for {
		fmt.Println("\n\nEscolha de qual UF deseja visualizar dados: ")
		fmt.Printf("Digite -1 para sair\n\n")
		for i := range states {
			fmt.Printf("%s ", states[i].Sigla)
		}
		fmt.Print("\n-> ")

		var UF string
		fmt.Scan(&UF)

		MostParticipantsUF(states)

		// verifica se existe UF no arrays de structs states
		if Contains(states, UF) == true {
			for i := range states {
				if UF == states[i].Sigla {
					PrintUFMeanScores(states[i])
					PrintUFRacesMeanScores(states[i])
				}
			}
		} else if UF == "-1" {
			break
		} else {
			fmt.Print("UF digitada inválida!\n")
			break
		}

		// reseta UF digitada pelo usuario
		UF = ""

		// tentativa de clear screen
		c := exec.Command("cls")
		c.Stdout = os.Stdout
		c.Run()
	}
}

//MenuRaces para escolher quais dados mostrar sobre raças
func MenuRaces(state State) {
	fmt.Printf("\n\nDeseja visualizar dados por raça nesse estado? (s/n)")
	fmt.Print("\n-> ")
	var opcao string
	fmt.Scan(&opcao)

	if opcao == "n"{
		return
	}
	for {
		fmt.Println("\nEscolha a raça.")
		fmt.Println("Digite -1 para sair.")
		fmt.Printf("1-Não declarada\n2-Branca\n3-Preta\n4-Parda\n5-Amarela\n6-Indígena\n")
		fmt.Print("\n-> ")

		var raca int
		fmt.Scan(&raca)

		if raca == -1{
			break
		}
		if raca > 6{
			fmt.Println("Opção inválida.")
			break
		}

		PrintUFRacesMeanScores(state)
	}
}

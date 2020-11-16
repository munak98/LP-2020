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

func MenuRaces(states State) {
	fmt.Printf("\n\nDeseja visualizar dados por raça nesse estado? (s/n)")
	fmt.Print("\n-> ")
	var opcao string
	fmt.Scan(&opcao)

	if opcao == "n"{
		return
	}
	for {
		fmt.Println("\nEscolha a raça.")
		fmt.Println("Digite 0 para sair.\n")
		fmt.Printf("1-Não declarada\n2-Branca\n3-Preta\n4-Parda\n5-Amarela\n6-Indígena\n")
		fmt.Print("\n-> ")

		var raca int
		fmt.Scan(&raca)

		if raca == 0{
			break
		}
		if raca > 6{
			fmt.Println("Opção inválida.")
		}
		PrintRacesMeanScores(states, raca-1)

	}
}


//Menu para escolher quais dados mostrar
func Menu(states []State) {
	fmt.Printf("\n\nDeseja visualizar dados de uma UF específica? (s/n)")
	fmt.Print("\n-> ")
	var opcao string
	fmt.Scan(&opcao)

	if opcao == "n"{
		return
	}

	for {
		fmt.Println("\nEscolha a UF: ")
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
					MenuRaces(states[i])
					fmt.Println("---------------------------------------")
					// PrintRacesMeanScores(states[i])
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

func YearsMenu(states19 []State, states18 []State, states17 []State, school *[3]SchoolScores, race *[3]RaceScores) {

	for {
		fmt.Println("Escolha o ano.")
		fmt.Printf("Digite 0 para sair\n\n")

		fmt.Printf("1-Enem 2017\n2-Enem 2018\n3-Enem 2019\n")

		fmt.Print("\n-> ")

		var year int
		fmt.Scan(&year)

		fmt.Printf("\n*****************************************\n")
		// verifica se existe UF no arrays de structs states
		switch year {
				case 0:
					break
				case 1:
					fmt.Printf("Dados nacionais Enem 2017\n")
					ShowGeneralRace((*race)[0])
					ShowGeneralSchools((*school)[0])
					Menu(states17)
				case 2:
					fmt.Printf("Dados nacionais Enem 2018\n")
					ShowGeneralRace((*race)[1])
					ShowGeneralSchools((*school)[1])
					Menu(states18)
				case 3:
					fmt.Printf("Dados nacionais Enem 2019\n")
					ShowGeneralRace((*race)[2])
					ShowGeneralSchools((*school)[2])
					Menu(states19)
				default:
					fmt.Print("Opção inválida!")
		}
		fmt.Printf("\n*****************************************\n")
	}
}


// pega o total de registros do arquivo CSV
func totalRecords(csvFilePath string) int {

	reader, _ := CsvReader(csvFilePath)

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

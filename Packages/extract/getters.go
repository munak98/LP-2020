package extract

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"sync"

	"github.com/montanaflynn/stats"
)

//GetData Pega os dados de forma paralela, lendo pedaços do arquivo
func getData(
	reader *csv.Reader,
	year *Year,
	start int,
	end int,
) {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		// leitura de pedaços do csv paralelamente
		for i := start; i < end; i++ {
			recordLine, err := reader.Read()

			if err == io.EOF {
				break // chegou ao final do registro
			} else if err != nil { //checa por outros erros
				fmt.Println("An error encountered ::", err)
			}

			for i := range (*year).States {
				if (*year).States[i].Sigla == recordLine[5] {

					(*year).States[i].Total++

					// Pega as notas nacionais de cada tipo de escola
					getYearSchoolScores(recordLine, &(*year))

					// Pega as notas estaduais de areas de conhecimento
					getScores(recordLine, &(*year).States[i].Scores, year.Year)

					// Coleta dados de cada raça
					getRacesData(recordLine, &(*year).States[i], year.Year)
				}
			}
		}
	}()
	wg.Wait()
}

// Pega os Scores nacionais de cada ano por tipo de escola do ensino medio
func getYearSchoolScores(recordLine []string, year *Year) {
	schoolType := getIntValue(recordLine, 17) // Tipo de Escola - Campo 17

	switch schoolType {
	case 1: // Nao respondeu
		getScores(recordLine, &year.SchoolScores[0], year.Year)
		year.SchoolType[schoolType-1]++
	case 2: // Publica
		getScores(recordLine, &year.SchoolScores[1], year.Year)
		year.SchoolType[schoolType-1]++
	case 3: // Privada
		getScores(recordLine, &year.SchoolScores[2], year.Year)
		year.SchoolType[schoolType-1]++
	case 4: // Exterior
		getScores(recordLine, &year.SchoolScores[3], year.Year)
		year.SchoolType[schoolType-1]++
	default:
		fmt.Println("Algo errado, tipo de escola invalido!")
	}
}

func getScores(recordLine []string, scores *[4][]float64, year int) {
	// Ano 2019
	if year == 2019 {
		scores[0] = append(scores[0], getScore(recordLine, 91)) // Ciencias da Natureza - campo 91
		scores[1] = append(scores[1], getScore(recordLine, 92)) // Ciencias Humanas - campo 92
		scores[2] = append(scores[2], getScore(recordLine, 93)) // Linguagens e Código - campo 93
		scores[3] = append(scores[3], getScore(recordLine, 94)) // Matemática - campo 94
	}

	// Anos 2017 e 2018
	scores[0] = append(scores[0], getScore(recordLine, 90)) // Ciencias da Natureza - campo 90
	scores[1] = append(scores[1], getScore(recordLine, 91)) // Ciencias Humanas - campo 91
	scores[2] = append(scores[2], getScore(recordLine, 92)) // Linguagens e Código - campo 92
	scores[3] = append(scores[3], getScore(recordLine, 93)) // Matemática - campo 93
}

func getScore(recordLine []string, campo int) float64 {
	i, _ := strconv.ParseFloat(recordLine[campo], 64)
	return i
}

func getIntValue(recordLine []string, campo int) int {
	i, _ := strconv.Atoi(recordLine[campo])
	return i
}

// Pega os dados de cada tipo de Raça
func getRacesData(recordLine []string, s *State, year int) {

	// Tipo de cor/raça - campo 9
	raceType := getIntValue(recordLine, 9)

	// Tipo de escola - campo 17
	schoolType := getIntValue(recordLine, 17)

	switch raceType {
	case 0: // Não informado
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	case 1: // Branca
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	case 2: // Preta
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	case 3: // Parda
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	case 4: // Amarela
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	case 5: // Indigena
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	default:
		fmt.Println("Raça não reconhecida, possível E.T!")
	}
}

// Pega os dados de uma Raça
func getRaceTypeData(recordLine []string, s *State, raceType int, schoolType int, year int) {
	s.Races[raceType].Total++
	getScores(recordLine, &s.Races[raceType].Scores, year) // Pega as notas de areas de conhecimento de Cada Raça
	getSchoolsData(recordLine, s, raceType, schoolType, year)
}

// pega os dados estaduais de cada tipo de escola e raça
func getSchoolsData(recordLine []string, s *State, raceType int, schoolType int, year int) {
	switch schoolType {
	case 1: // Nao respondeu
		getSchoolTypeData(recordLine, s, raceType, schoolType-1, year)
	case 2: // Publica
		getSchoolTypeData(recordLine, s, raceType, schoolType-1, year)
	case 3: // Privada
		getSchoolTypeData(recordLine, s, raceType, schoolType-1, year)
	case 4: // Exterior
		getSchoolTypeData(recordLine, s, raceType, schoolType-1, year)
	default:
		fmt.Println("Algo errado, tipo de escola invalido!")
	}
}

func getSchoolTypeData(recordLine []string, s *State, raceType int, schoolType int, year int) {
	s.SchoolType[schoolType]++
	s.Races[raceType].SchoolType[schoolType]++
	getScores(recordLine, &s.SchoolScores[schoolType], year)
	getScores(recordLine, &s.Races[raceType].SchoolScores[schoolType], year)
}

// Calcula as medias das notas de cada area de conhecimento
func getMeanScores(scores [4][]float64) [4]float64 {
	meanScores := [4]float64{}
	for i := range scores {
		meanScores[i], _ = stats.Mean(scores[i])
	}
	return meanScores
}

// Pega as medias estaduais
func getStatesMeanScores(states *[]State) {
	for i := range *states {
		// Pega as medias estaduais
		(*states)[i].Medias = getMeanScores((*states)[i].Scores)

		// Pega as medias estaduais de cada tipo de escola
		for j := range (*states)[i].SchoolMeanScores {
			(*states)[i].SchoolMeanScores[j] = getMeanScores((*states)[i].SchoolScores[j])
		}

		for j := range (*states)[i].Races {
			// Pega as medias estaduais de cada raça
			(*states)[i].Races[j].Medias = getMeanScores((*states)[i].Races[j].Scores)

			// Pega as medias estaduais de cada tipo de escola de cada raça
			for k := range (*states)[i].Races[j].SchoolMeanScores {
				(*states)[i].Races[j].SchoolMeanScores[k] =
					getMeanScores((*states)[i].Races[j].SchoolScores[k])
			}
		}
	}
}

// Pega as medias nacionais
func getYearsMeanScores(years *[]Year) {
	statesMeanScores := [4][]float64{}
	for i := range *years {
		for j := range (*years)[i].States {

			// pega as medias de cada estado
			for k := range (*years)[i].States[j].Medias {
				statesMeanScores[k] = append(statesMeanScores[k], (*years)[i].States[j].Medias[k])
			}

			// Pega o total nacional de participantes de cada raça
			for k := range (*years)[i].States[j].Races {
				(*years)[i].Races[k].Total += (*years)[i].States[j].Races[k].Total

				// pega as medias estaduais das raças
				for l := range (*years)[i].States[j].Races[k].Medias {
					(*years)[i].Races[k].Scores[l] =
						append((*years)[i].Races[k].Scores[l], (*years)[i].States[j].Races[k].Medias[l])
				}
			}
		}
		// pega as medias nacionais
		(*years)[i].Medias = getMeanScores(statesMeanScores)

		// pega as medias nacionais de cada tipo de escola
		for j := range (*years)[i].SchoolMeanScores {
			(*years)[i].SchoolMeanScores[j] = getMeanScores((*years)[i].SchoolScores[j])
		}

		// tira as medias nacionais das medias estaduais de cada raça
		for k := range (*years)[i].Races {
			(*years)[i].Races[k].Medias = getMeanScores((*years)[i].Races[k].Scores)
		}
	}
}

package concurrency_extract

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"sync"

	"github.com/montanaflynn/stats"
)

func getScore(recordLine []string, campo int) float64 {
	i, _ := strconv.ParseFloat(recordLine[campo], 64)
	return i
}

func getIntValue(recordLine []string, campo int) int {
	i, _ := strconv.Atoi(recordLine[campo])
	return i
}

func getScores(recordLine []string, scores [4][]float64, year int) [4][]float64 {
	if year == 2{
		scores[0] = append(scores[0], getScore(recordLine, 91)) // Ciencias da Natureza - campo 91
		scores[1] = append(scores[1], getScore(recordLine, 92)) // Ciencias Humanas - campo 92
		scores[2] = append(scores[2], getScore(recordLine, 93)) // Linguagens e Código - campo 93
		scores[3] = append(scores[3], getScore(recordLine, 94)) // Matemática - campo 94
	}	else{
		scores[0] = append(scores[0], getScore(recordLine, 90)) // Ciencias da Natureza - campo 90
		scores[1] = append(scores[1], getScore(recordLine, 91)) // Ciencias Humanas - campo 91
		scores[2] = append(scores[2], getScore(recordLine, 92)) // Linguagens e Código - campo 92
		scores[3] = append(scores[3], getScore(recordLine, 93)) // Matemática - campo 93
	}

	return scores
}

// Pega as notas de areas de conhecimento de cada UF
func getUFScores(recordLine []string, state *State, year int, wg *sync.WaitGroup) {
	defer wg.Done()
	state.Scores = getScores(recordLine, state.Scores, year)
	return
}

// Pega as notas de areas de conhecimento de Cada Raça
func getRaceScores(recordLine []string, state *State, raceType int, year int) {
	state.Races[raceType].Scores = getScores(recordLine, state.Races[raceType].Scores, year)
	return
}

// Pega as medias das notas de cada area de conhecimento
func getMeanScores(scores [4][]float64) [4]float64 {
	meanScores := [4]float64{}

	for i := range scores {
		meanScores[i], _ = stats.Mean(scores[i])
	}

	return meanScores
}

// Pega as medias das notas de cada area de conhecimento e por raça de todos Estados
func getStatesMeanScores(states *[]State) {
	for i := range *states {
		(*states)[i].Medias = getMeanScores((*states)[i].Scores)

		for j := range (*states)[i].Races {
			(*states)[i].Races[j].Medias = getMeanScores((*states)[i].Races[j].Scores)
		}
	}
	return
}

func getSchoolScores(recordLine []string, scores *SchoolScores, year int, wg *sync.WaitGroup) {
	defer wg.Done()
	if getIntValue(recordLine, 17) == 1{
		scores.Public = getScores(recordLine, scores.Public, year)
	}
	if getIntValue(recordLine, 17) == 2{
		scores.Private = getScores(recordLine, scores.Private, year)
	}
	return
}

func getGeneralRaceScores(recordLine []string, scores *RaceScores, year int, wg *sync.WaitGroup) {
	defer wg.Done()
	raceType := getIntValue(recordLine, 9)
	switch raceType {
	case 0: // Não informado
		scores.NaoDeclarada = getScores(recordLine, scores.NaoDeclarada, year)
		break
	case 1: // Branca
		scores.BrcPrdAmar = getScores(recordLine, scores.BrcPrdAmar, year)
		break
	case 2: // Preta
		scores.Preta = getScores(recordLine, scores.Preta, year)
		break
	case 3: // Parda
		scores.BrcPrdAmar = getScores(recordLine, scores.BrcPrdAmar, year)
		break
	case 4: // Amarela
		scores.BrcPrdAmar = getScores(recordLine, scores.BrcPrdAmar, year)
		break
	case 5: // Indigena
		scores.Indigena = getScores(recordLine, scores.Indigena, year)
		break
	default:
		fmt.Println("Raça não reconhecida.")
		break
	}

	return
}

// Pega os dados de cada tipo de Raça
func getRacesData(recordLine []string, s *State, year int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Tipo de cor/raça - campo 9
	raceType := getIntValue(recordLine, 9)

	// Tipo de escola - campo 17
	schoolType := getIntValue(recordLine, 17)

	switch raceType {
	case 0: // Não informado
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
		break
	case 1: // Branca
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
		break
	case 2: // Preta
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
		break
	case 3: // Parda
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
		break
	case 4: // Amarela
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
		break
	case 5: // Indigena
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
		break
	default:
		fmt.Println("Raça não reconhecida, possível E.T!")
		break
	}

	return
}

// Pega os dados de uma Raça
func getRaceTypeData(recordLine []string, s *State, raceType int, schoolType int, year int) {
	s.Races[raceType].Total++
	getRaceScores(recordLine, s, raceType, year)
	getSchoolType(s, raceType, schoolType)
}

// pega os dados para cada tipo de escola
func getSchoolType(s *State, raceType int, schoolType int) {

	switch schoolType {
	case 1: // Nao respondeu
		s.SchoolType[0]++
		s.Races[raceType].SchoolType[0]++
		break
	case 2: // Publica
		s.SchoolType[1]++
		s.Races[raceType].SchoolType[1]++
		break
	case 3: // Privada
		s.SchoolType[2]++
		s.Races[raceType].SchoolType[2]++
		break
	case 4: // Exterior
		s.SchoolType[3]++
		s.Races[raceType].SchoolType[3]++
		break
	default:
		fmt.Println("Algo errado, tipo de escola Inválido!")
		fmt.Println("%d", schoolType)
		break
	}
}

//GetData Pega os dados de forma paralela, lendo pedaços do arquivo
func getData(
	reader *csv.Reader,
	states *[]State,
	schoolScores *SchoolScores,
	raceScores *RaceScores,
	count *int,
	begin int,
	end int,
	year int,
) {

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()

		// leitura de pedaços do csv paralelamente
		for i := begin; i < end; i++ {
			recordLine, err := reader.Read()

			if err == io.EOF {
				break // chegou ao final do registro
			} else if err != nil { //checa por outros erros
				fmt.Println("An error encountered ::", err)
			}
			*count++

			getGeneralRaceScores(recordLine, raceScores, year, wg)

			getSchoolScores(recordLine, schoolScores, year, wg)

			for i := range *states {
				if (*states)[i].Sigla == recordLine[5] {

					(*states)[i].Total++

					// coleta as notas de cada area da UF
					getUFScores(recordLine, &(*states)[i], year, wg)

					// coleta dados de cada area por raça da UF
					getRacesData(recordLine, &(*states)[i], year, wg)
				}
			}
		}
	}()

	wg.Wait()
}
